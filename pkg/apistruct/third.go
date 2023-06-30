// Copyright © 2023 OpenIM. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package apistruct

import "mime/multipart"

type MinioStorageCredentialReq struct {
	OperationID string `json:"operationID"`
}

type MiniostorageCredentialResp struct {
	SecretAccessKey  string `json:"secretAccessKey"`
	AccessKeyID      string `json:"accessKeyID"`
	SessionToken     string `json:"sessionToken"`
	BucketName       string `json:"bucketName"`
	StsEndpointURL   string `json:"stsEndpointURL"`
	StorageTime      int    `json:"storageTime"`
	IsDistributedMod bool   `json:"isDistributedMod"`
}

type MinioUploadFileReq struct {
	OperationID string `form:"operationID" binding:"required"`
	FileType    int    `form:"fileType" binding:"required"`
}

type MinioUploadFile struct {
	URL             string `json:"URL"`
	NewName         string `json:"newName"`
	SnapshotURL     string `json:"snapshotURL,omitempty"`
	SnapshotNewName string `json:"snapshotName,omitempty"`
}

type MinioUploadFileResp struct {
	Data struct {
		MinioUploadFile
	} `json:"data"`
}

type UploadUpdateAppReq struct {
	OperationID string                `form:"operationID" binding:"required"`
	Type        int                   `form:"type" binding:"required"`
	Version     string                `form:"version"  binding:"required"`
	File        *multipart.FileHeader `form:"file" binding:"required"`
	Yaml        *multipart.FileHeader `form:"yaml"`
	ForceUpdate bool                  `form:"forceUpdate"`
	UpdateLog   string                `form:"updateLog" binding:"required"`
}

type UploadUpdateAppResp struct {
}

type GetDownloadURLReq struct {
	OperationID string `json:"operationID" binding:"required"`
	Type        int    `json:"type" binding:"required"`
	Version     string `json:"version" binding:"required"`
}

type GetDownloadURLResp struct {
	Data struct {
		HasNewVersion bool   `json:"hasNewVersion"`
		ForceUpdate   bool   `json:"forceUpdate"`
		FileURL       string `json:"fileURL"`
		YamlURL       string `json:"yamlURL"`
		Version       string `json:"version"`
		UpdateLog     string `json:"update_log"`
	} `json:"data"`
}

type GetRTCInvitationInfoReq struct {
	OperationID string `json:"operationID" binding:"required"`
	ClientMsgID string `json:"clientMsgID" binding:"required"`
}

type GetRTCInvitationInfoResp struct {
	Data struct {
		OpUserID   string `json:"opUserID"`
		Invitation struct {
			InviterUserID     string   `json:"inviterUserID"`
			InviteeUserIDList []string `json:"inviteeUserIDList"`
			GroupID           string   `json:"groupID"`
			RoomID            string   `json:"roomID"`
			Timeout           int32    `json:"timeout"`
			MediaType         string   `json:"mediaType"`
			SessionType       int32    `json:"sessionType"`
			InitiateTime      int32    `json:"initiateTime"`
			PlatformID        int32    `json:"platformID"`
			CustomData        string   `json:"customData"`
		} `json:"invitation"`
		OfflinePushInfo struct{} `json:"offlinePushInfo"`
	} `json:"data"`
}

type GetRTCInvitationInfoStartAppReq struct {
	OperationID string `json:"operationID" binding:"required"`
}

type GetRTCInvitationInfoStartAppResp struct {
	GetRTCInvitationInfoResp
}

/**
 * FCM第三方上报Token
 */
type FcmUpdateTokenReq struct {
	OperationID string `json:"operationID" binding:"required"`
	Platform    int    `json:"platform" binding:"required,min=1,max=2"` //only for ios + android
	FcmToken    string `json:"fcmToken" binding:"required"`
}

type FcmUpdateTokenResp struct {
}
type SetAppBadgeReq struct {
	OperationID    string `json:"operationID" binding:"required"`
	FromUserID     string `json:"fromUserID" binding:"required"`
	AppUnreadCount int32  `json:"appUnreadCount"`
}

type SetAppBadgeResp struct {
}
