package cmd

import (
	"fmt"
	"path"
	"time"

	"github.com/Rxg1898/ProductsStorage/store"
	"github.com/Rxg1898/ProductsStorage/store/provider/aliyun"
	"github.com/spf13/cobra"
)

const (
	defaultBuckName = ""
	defaultEndpoint = ""
	defaultALIAK    = ""
	defaultALISK    = ""
)

var (
	buckName       string
	uploadFilePath string
	bucketEndpoint string
)

// uploadCmd represents the start command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "上传文件到制品库",
	Long:  `上传文件到制品库`,
	RunE: func(cmd *cobra.Command, args []string) error {
		p, err := getProvider()
		if err != nil {
			return err
		}
		if uploadFilePath == "" {
			return fmt.Errorf("upload file path is missing")
		}
		day := time.Now().Format("20060102")
		fn := path.Base(uploadFilePath)
		ok := fmt.Sprintf("%s/%s", day, fn)
		err = p.UploadFile(buckName, ok, uploadFilePath)
		if err != nil {
			return err
		}
		return nil
	},
}

func getProvider() (store.Uploader, error) {
	switch ossProvider {
	case "aliyun":
		return aliyun.NewUploader(bucketEndpoint, aliAccessID, aliAccessKey)
	case "qccloud":
		return nil, fmt.Errorf("not impl")
	case "minio":
		return nil, fmt.Errorf("not impl")
	default:
		return nil, fmt.Errorf("unknown uploader %s", ossProvider)

	}
}

func init() {
	uploadCmd.PersistentFlags().StringVarP(&uploadFilePath, "file_path", "f", "", "upload file path")
	uploadCmd.PersistentFlags().StringVarP(&buckName, "bucket_name", "b", defaultBuckName, "upload oss bucket name")
	uploadCmd.PersistentFlags().StringVarP(&bucketEndpoint, "bucket_endpoint", "e", defaultEndpoint, "upload oss endpoint")
	RootCmd.AddCommand(uploadCmd)
}
