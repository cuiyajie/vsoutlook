package site

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"vsoutlook.com/vsoutlook/infra/config"
	"vsoutlook.com/vsoutlook/infra/redis"
)

type deploymentInfo struct {
	Commit string `json:"commit"`
	Index  string `json:"index"`
}

func startSiteApp() bool {
	return true
}

// siteArea: cn, intl
func loadForSite(region string, branch string) (err error) {
	// https://static.bench3d.com/intl/latest-86e5dda730a2f7c.txt
	now := time.Now().Unix()
	var infoFileName = latestInfoFileName
	if strings.Contains(config.ClusterID, "prod") {
		infoFileName = prodInfoFileName
	} else {
		if branch != "main" && branch != "" {
			infoFileName = strings.ReplaceAll(infoFileName, "latest", branch)
		}
	}
	url := fmt.Sprintf("%s/%s/%s?t=%d", staticSiteURL, region, infoFileName, now)
	info := deploymentInfo{}
	if err = parseJSONFromURL(url, &info); err != nil {
		return err
	}

	// site-cn, site-intl
	var dir = RegionDir(info.Region)
	var prefix = staticSiteURL + "/" + info.Region + "/"
	var pairs = [][]string{
		{prefix + info.Index, "app"},
		// {prefix + info.Capture, "capture"},
	}

	if config.EnableStorybook && region == def.CN {
		pairs = append(pairs, []string{
			staticSiteURL + "/storybook/" + info.Storybook, "storybook"})
	}

	for _, pair := range pairs {
		target := filepath.Join(dir, pair[1]+".html")
		if err = download(pair[0], target); err != nil {
			return err
		}

		key := region + "/" + pair[1]
		appHTMLStore.Store(key, readFile(target))
	}
	redis.Set("frontend-commit", info.Commit)
	return err
}
