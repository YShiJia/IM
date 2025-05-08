/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2025-04-29 12:01:10
 */

package init

import log "github.com/sirupsen/logrus"

func InitLog() error {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	log.SetReportCaller(true)
	log.SetLevel(log.DebugLevel)
	return nil
}
