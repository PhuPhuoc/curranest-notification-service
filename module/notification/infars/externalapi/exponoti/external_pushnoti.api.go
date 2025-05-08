package exponoti

import (
	"context"
	"encoding/json"
	"fmt"

	"githuh.com/PhuPhuoc/curanest-notification-service/common"
	notificationcommands "githuh.com/PhuPhuoc/curanest-notification-service/module/notification/usecase/commands"
)

func (api *externalExpoNotiService) PushNoti(ctx context.Context, req *notificationcommands.PushNotification) error {
	result, err := common.CallExternalAPI(ctx, common.RequestOptions{
		Method:  "POST",
		URL:     api.apiURL,
		Payload: req,
	})
	if err != nil {
		return common.NewInternalServerError().WithReason("cannot call external api: " + err.Error())
	}

	data, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("failed to marshal result: %v", err)
	}

	var apiResponse notificationcommands.ExpoResponse
	if err := json.Unmarshal(data, &apiResponse); err != nil {
		return fmt.Errorf("failed to parse JSON: %v, body: %s", err, string(data))
	}

	if apiResponse.Data.Status != "ok" {
		return fmt.Errorf("API returned non-OK status '%s'", apiResponse.Data.Status)
	}

	return nil
}
