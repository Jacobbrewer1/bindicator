package config

type StructConfig struct {
	RemoteConfig      *RemoteConfigStruct
	ConnectionStrings *ConnectionStringsStruct `json:"ConnectionStrings,omitempty"`
	JsonConfig        *JsonConfigStruct        `json:"JsonConfig,omitempty"`
}

type RemoteConfigStruct struct {
	Email      *EmailConfigStruct `json:"Email"`
	ApiSecrets *ApiSecretsStruct  `json:"ApiSecrets,omitempty"`
	People     []*PeopleConfig    `json:"People,omitempty"`
}

type EmailConfigStruct struct {
	From     *string `json:"From"`
	Password *string `json:"Password"`
	SmtpHost *string `json:"SmtpHost"`
	SmtpPort *string `json:"SmtpPort"`
}

type PeopleConfig struct {
	Name  *string `json:"Name,omitempty"`
	Email *string `json:"Email,omitempty"`
	UPRN  *string `json:"UPRN,omitempty"`
	Bins  `json:"Bins,omitempty"`
}

type JsonConfigStruct struct {
	ConfigIpAddress *string `json:"ConfigIpAddress"`
}

type ApiSecretsStruct struct {
	GiphyApiToken *string `json:"GiphyApiToken,omitempty"`
}

type ConnectionStringsStruct struct {
	BCPCouncil  *string `json:"BCPCouncil,omitempty"`
	OpenWeather *string `json:"OpenWeather,omitempty"`
	NewsApi     *string `json:"NewsApi,omitempty"`
}
