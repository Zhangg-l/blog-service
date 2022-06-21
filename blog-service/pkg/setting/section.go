package setting

import "time"

var sections = make(map[string]interface{})

type ServerSetting struct {
	RunMode      string
	HttpPort     string
	Readtimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSetting struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
	// 上传图片
	UploadImageMaxSize   int
	UploadSavePath       string
	UploadServerUrl      string
	UploadImageAllowExts []string

	// 执行时间
	DefaultContextTimeout time.Duration
}

type DatabaseSetting struct {
	DBType       string
	Password     string
	Username     string
	DBName       string
	Tableprefix  string
	Charset      string
	Host         string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type JWTSetting struct {
	Secret string
	Expire time.Duration
	Issuer string
}

type EmailSetting struct {
	Host     string
	Port     int
	UserName string
	IsSSL    bool
	From     string
	To       []string
	Password string
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)

	if err != nil {
		return err
	}
	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

func (s *Setting) ReloadAllSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
