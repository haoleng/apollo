//Copyright (c) 2017 Phil

//Package apollo ctrip apollo go client
package apollo

var (
	defaultClient *Client
)

// Start apollo
func Start() error {
	return StartWithConfFile(defaultConfName)
}

// StartWithConfFile run apollo with conf file
func StartWithConfFile(name string) error {
	conf, err := NewConf(name)
	if err != nil {
		return err
	}
	return StartWithConf(conf)
}

// StartWithConf run apollo with Conf
func StartWithConf(conf *Conf) error {
	defaultClient = NewClient(conf)

	return defaultClient.Start()
}

// Stop sync config
func Stop() error {
	return defaultClient.Stop()
}

// WatchUpdate get all updates
func WatchUpdate() <-chan *ChangeEvent {
	return defaultClient.WatchUpdate()
}

// GetStringValueWithNameSapce get value from given namespace
func GetStringValueWithNameSapce(namespace, key, defaultValue string) string {
	return defaultClient.GetStringValueWithNameSapce(namespace, key, defaultValue)
}

// GetStringValue from default namespace
func GetStringValue(key, defaultValue string) string {
	return GetStringValueWithNameSapce(defaultNamespace, key, defaultValue)
}

// GetNameSpaceContent get contents of namespace
func GetNameSpaceContent(namespace, defaultValue string) string {
	return defaultClient.GetNameSpaceContent(namespace, defaultValue)
}