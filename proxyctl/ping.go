package proxyctl

// func ParallelMeasureDelay(links []*subscription.Link, threads int, times int, timeout uint64) []*subscription.Link {
// 	if threads <= 0 {
// 		threads = 1
// 	}
// 	slog.Debug("ping with %d threads", threads)

// 	threadPool := pool.New().WithMaxGoroutines(threads)
// 	for _, link := range links {
// 		link0 := link // capture range variable

// 		threadPool.Go(func() {
// 			port := PickFreeTcpPort()
// 			delay, err := MeasureDelay(link0, port, times, timeout)
// 			if err != nil {
// 				log.Error(err)
// 				delay = -1
// 			}

// 			link0.Delay = delay
// 		})
// 	}
// 	threadPool.Wait()

// 	var r []*subscription.Link
// 	for _, lk := range links {
// 		if lk.Delay > -1 {
// 			r = append(r, lk)
// 		}
// 	}

// 	if len(r) == 0 {
// 		return r
// 	}
// 	return sortByDelay(r)
// }

// func MeasureDelay(link *subscription.Link, port int, times int, timeout uint64) (int32, error) {
// 	a, err := NewXray([]*subscription.Link{link}, []*settings.Listen{listen}, false, false, -1)
// 	if err != nil {
// 		return -1, err
// 	}

// 	if err := a.Start(); err != nil {
// 		return -1, err
// 	}
// 	defer a.Close()

// 	client := a.NewHttpClient(timeout)
// 	return ping(client, times, link.Remarks), nil
// }

// func sortByDelay(links []*subscription.Link) []*subscription.Link {
// 	sort.Slice(links, func(a, b int) bool {
// 		return links[a].Delay < links[b].Delay
// 	})
// 	return links
// }
