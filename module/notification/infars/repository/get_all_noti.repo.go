package notificationrepository

import (
	"context"
	"fmt"
	"math"
	"strings"
	"sync"

	"githuh.com/PhuPhuoc/curanest-notification-service/common"
	notificationdomain "githuh.com/PhuPhuoc/curanest-notification-service/module/notification/domain"
	notificationqueries "githuh.com/PhuPhuoc/curanest-notification-service/module/notification/usecase/quries"
)

func (repo *notificationRepo) GetAllNotification(ctx context.Context, filter *notificationqueries.FilterGetNoti) ([]notificationdomain.Notification, error) {
	filter.Paging.Process()
	var whereConditions []string
	var args []interface{}

	if filter.AccountID.String() != "" {
		whereConditions = append(whereConditions, "account_id = ?")
		args = append(args, filter.AccountID.String())
	}

	if filter.HadRead != nil {
		if *filter.HadRead {
			whereConditions = append(whereConditions, "read_at is not null")
		} else {
			whereConditions = append(whereConditions, "read_at is null")
		}
	}

	orderBy := " order by created_at desc "

	limit := ""

	var argsLimit []interface{}
	limit = " limit ? offset ?"
	argsLimit = append(argsLimit, filter.Paging.Size, (filter.Paging.Page-1)*filter.Paging.Size)

	var where string
	if len(whereConditions) > 0 {
		where = strings.Join(whereConditions, " AND ")
	}

	queryGetData := common.GenerateSQLQueries(common.SELECT_WITHOUT_COUNT, TABLE_NOTIFICATION, GET_NOTIFICATION, &where)
	queryGetData += orderBy + limit
	queryGetCount := common.GenerateSQLQueries(common.SELECT_COUNT, TABLE_NOTIFICATION, GET_NOTIFICATION, &where)

	errchan := make(chan error, 2)
	countchan := make(chan int, 1)
	datachan := make(chan []notificationdomain.Notification, 1)
	argsOfGetData := append(args, argsLimit...)

	var wg sync.WaitGroup
	wg.Add(2)
	go repo.getNotificationCount(ctx, queryGetCount, args, errchan, countchan, &wg)
	go repo.getNotificationData(ctx, queryGetData, argsOfGetData, errchan, datachan, &wg)

	var once sync.Once
	go func() {
		defer func() {
			if r := recover(); r != nil {
				errchan <- fmt.Errorf("panic in goroutine: %v", r)
			}
		}()
		wg.Wait()
		once.Do(func() {
			close(errchan)
			close(countchan)
			close(datachan)
		})
	}()

	var totalRecord int
	var notifications []notificationdomain.Notification

	receivedCount := 0
	expectedCount := 2

	for {
		select {
		case err, ok := <-errchan:
			if ok {
				return nil, err
			}
		case count, ok := <-countchan:
			if ok {
				totalRecord = count
				receivedCount++
			}
		case data, ok := <-datachan:
			if ok {
				notifications = data
				receivedCount++
			}
		case <-ctx.Done():
			return nil, fmt.Errorf("operation timed out: %w", ctx.Err())
		}

		if receivedCount == expectedCount {
			break
		}
	}

	totalPages := int(math.Ceil(float64(totalRecord) / float64(filter.Paging.Size)))
	filter.Paging.Total = totalPages
	return notifications, nil
}

func (repo *notificationRepo) getNotificationData(
	ctx context.Context,
	queryStr string,
	args []interface{},
	errchan chan<- error,
	datachan chan<- []notificationdomain.Notification,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	var dtos []NotificationDTO
	if err := repo.db.SelectContext(ctx, &dtos, queryStr, args...); err != nil {
		errchan <- err
		return
	}

	entities := make([]notificationdomain.Notification, len(dtos))
	for i := range dtos {
		entity, _ := dtos[i].ToNotificationEntity()
		entities[i] = *entity
	}

	datachan <- entities
}

func (repo *notificationRepo) getNotificationCount(
	ctx context.Context,
	queryStr string,
	args []interface{},
	errchan chan<- error,
	countchan chan<- int,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	var total int
	if err := repo.db.GetContext(ctx, &total, queryStr, args...); err != nil {
		errchan <- err
		return
	}

	countchan <- total
}
