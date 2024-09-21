//	@title			Serein Shop server
//	@version		1.0
//	@description	This is serein shop server v1.0
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Serein shop support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	jiahuipaung@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:5001
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

package main

import (
	"fmt"

	conf "serein/config"
	util "serein/pkg/utils/log"
	_ "serein/docs"

	// "serein/pkg/utils/track"
	// "serein/repository/cache"
	"serein/repository/db/dao"
	// "serein/repository/es"
	// "serein/repository/kafka"
	// "serein/repository/rabbitmq"
	"serein/routes"
)

func main() {
	loading() // 加载配置
	r := routes.NewRouter()
	_ = r.RunTLS(conf.Config.System.HttpPort,"server.crt", "server.key")
	fmt.Println("启动成功")
}

func loading() {
	conf.InitConfig()
	dao.InitMySQL()
	// cache.InitCache()
	// rabbitmq.InitRabbitMQ()
	// es.InitES()
	// kafka.InitKafka()
	// track.InitJaeger()
	util.InitLog()
	fmt.Println("配置加载完成...")
	go scriptStarting()
}

func scriptStarting() {

}
