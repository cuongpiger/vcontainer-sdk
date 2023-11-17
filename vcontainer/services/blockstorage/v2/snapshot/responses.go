package snapshot

import "github.com/vngcloud/vcontainer-sdk/vcontainer/services/blockstorage/v2/snapshot/obj"

type (
	CreateResponse struct {
		BackendID           string         `json:"backendId"`
		BackendSnapshotID   *string        `json:"backendSnapshotId"`
		BackendSnapshotName *string        `json:"backendSnapshotName"`
		CreatedAt           string         `json:"createdAt"`
		DeletedAt           *string        `json:"deletedAt"`
		ID                  string         `json:"id"`
		ParentID            *string        `json:"parentId"`
		ParentType          string         `json:"parentType"`
		ProjectID           string         `json:"projectId"`
		Size                int64          `json:"size"`
		SnapshotConfig      Config         `json:"snapshotConfig"`
		SnapshotTime        int64          `json:"snapshotTime"`
		SnapshotVolumeId    string         `json:"snapshotVolumeId"`
		Status              string         `json:"status"`
		UpdatedAt           string         `json:"updatedAt"`
		UserId              int64          `json:"userId"`
		VolumeId            string         `json:"volumeId"`
		VolumeSnapshot      VolumeSnapshot `json:"volumeSnapshot"`
		Name                string         `json:"name"`
		Description         string         `json:"description"`
	}

	Config struct {
		IsPermanently bool  `json:"isPermanently"`
		RetainedDays  int64 `json:"retainedDays"`
	}

	VolumeSnapshot struct {
		BackendPool      string  `json:"backendPool"`
		BackendPrefix    string  `json:"backendPrefix"`
		BackendStatus    string  `json:"backendStatus"`
		BackendUuid      string  `json:"backendUuid"`
		BootIndex        int     `json:"bootIndex"`
		Bootable         bool    `json:"bootable"`
		CreatedAt        string  `json:"createdAt"`
		Id               string  `json:"id"`
		MultiAttach      bool    `json:"multiAttach"`
		Name             string  `json:"name"`
		Product          string  `json:"product"`
		ProjectId        string  `json:"projectId"`
		Size             int64   `json:"size"`
		Status           string  `json:"status"`
		Type             string  `json:"type"`
		UpdatedAt        string  `json:"updatedAt"`
		Uuid             string  `json:"uuid"`
		VolumeId         int64   `json:"volumeId"`
		VolumeTypeId     string  `json:"volumeTypeId"`
		VolumeTypeZoneId string  `json:"volumeTypeZoneId"`
		EncryptionType   *string `json:"encryptionType"`
	}
)

func (s *CreateResponse) ToSnapshotObject() *obj.Snapshot {
	return &obj.Snapshot{
		ID:        s.ID,
		CreatedAt: s.CreatedAt,
		VolumeID:  s.VolumeId,
		Status:    s.Status,
		Size:      s.Size,
	}
}
