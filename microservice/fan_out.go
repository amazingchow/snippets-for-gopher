package microservice

// FanOut 一种信号通知模式,
// 常用于观察者模式中, 上游/下游服务状态变动后, 多个观察者都会接收到变更通知信号.
func FanOut(in <-chan interface{}, out []chan interface{}, async bool) {
	go func() {
		defer func() { // 退出时关闭所有的输出channel
			for i := 0; i < len(out); i++ {
				close(out[i])
			}
		}()

		for v := range in {
			v := v
			for i := 0; i < len(out); i++ {
				i := i
				// TODO: 异步模式下, 会出现向已关闭的channel写数据
				if async {
					go func() {
						out[i] <- v
					}()
				} else {
					out[i] <- v
				}
			}
		}
	}()
}