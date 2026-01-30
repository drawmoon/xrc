package subscription

// // GetSubscription fetches and parses subscription links from the given URL.
// func GetSubscription(url string) ([]Link, error) {
// 	req, _ := http.NewRequest("GET", url, nil)
// 	client := &http.Client{Timeout: 10 * time.Second}

// 	resp, err := client.Do(req)
// 	if err != nil || resp.StatusCode != 200 {
// 		return nil, fmt.Errorf("fetch subscription failed, url: %s", url)
// 	}
// 	defer resp.Body.Close()

// 	b, _ := io.ReadAll(resp.Body)
// 	links, err := parseSubscription(string(b))
// 	if err != nil {
// 		return nil, err
// 	}

// 	slog.Debug("found subscriptions", "count", len(links))
// 	return links, nil
// }

// // parseSubscription decodes a base64-encoded subscription string and parses it into Link objects.
// func parseSubscription(s string) ([]Link, error) {
// 	b, err := base64.StdEncoding.DecodeString(s)
// 	if err != nil {
// 		return nil, errors.New("decode subscription failed")
// 	}

// 	var links []Link

// 	lines := strings.Split(string(b), "\n")
// 	for _, lineStr := range lines {
// 		if len(lineStr) == 0 {
// 			continue
// 		}
// 		link, err := PaseLink(lineStr)
// 		if err != nil {
// 			slog.Warn("parse link failed", "link", lineStr, "error", err)
// 			continue
// 		}
// 		links = append(links, link)
// 	}

// 	return links, nil
// }
