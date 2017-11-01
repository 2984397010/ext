package command

import (
	"os"

	"github.com/gogo/protobuf/proto"
	"v2ray.com/core/common"
	"v2ray.com/ext/tools/conf/serial"
	"v2ray.com/ext/tools/control"
)

func init() {
	const name = "config"
	common.Must(control.RegisterCommand(name, func(args []string) {
		pbConfig, err := serial.LoadJSONConfig(os.Stdin)
		if err != nil {
			os.Stderr.WriteString("failed to parse json config: " + err.Error())
			return
		}

		bytesConfig, err := proto.Marshal(pbConfig)
		if err != nil {
			os.Stderr.WriteString("failed to marshal proto config: " + err.Error())
			return
		}

		if _, err := os.Stdout.Write(bytesConfig); err != nil {
			os.Stderr.WriteString("failed to write proto config: " + err.Error())
			return
		}
	}))
}
