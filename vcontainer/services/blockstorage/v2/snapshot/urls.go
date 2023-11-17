package snapshot

import "github.com/vngcloud/vcontainer-sdk/client"

// createURL generates the URL for creating a Snapshot based on the given options.
func createURL(pSc *client.ServiceClient, pOpts ICreateOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"volumes",
		pOpts.GetVolumeID(),
		"snapshots")
}

func deleteURL(pSc *client.ServiceClient, pOpts IDeleteOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"volumes",
		pOpts.GetVolumeID(),
		"snapshots",
		pOpts.GetSnapshotID(),
	)
}
