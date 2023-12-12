package config

import (
	"encoding"
	"errors"
	"reflect"
	"strings"
	
	"github.com/spf13/viper"
)

var Env = envConf{}

type envConf struct {
	Debug              bool   `env:"debug"`
	DBProvider         string `env:"database.provider"`
	DBHost             string `env:"database.host"`
	DBPort             int    `env:"database.port"`
	DBUser             string `env:"database.user"`
	DBPasswd           string `env:"database.pass"`
	DBName             string `env:"database.name"`
	CtxTimeout         int    `env:"context.timeout"`
	ServerAddr         string `env:"server.address"`
	ServerHostGRPC     string `env:"server.hostgrpc"`
	ServerAddrGRPC     string `env:"server.addressgrpc"`
	Timezone           string `env:"timezone"`
	PaymentApiHostGRPC string `env:"payment_api.host_grpc"`
}

func LoadEnv() {
	viper.SetConfigFile(`env.json`)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	
	if err := parse(&Env); err != nil {
		panic(err)
	}
}

//goland:noinspection GoErrorStringFormat
var (
	ErrNotAStructPtr   = errors.New("Expected a pointer to a Struct")
	ErrUnsupportedType = errors.New("Type is not supported")
)

func parse(v interface{}) error {
	ptrRef := reflect.ValueOf(v)
	if ptrRef.Kind() != reflect.Ptr {
		return ErrNotAStructPtr
	}
	ref := ptrRef.Elem()
	if ref.Kind() != reflect.Struct {
		return ErrNotAStructPtr
	}
	return doParse(ref)
}

func doParse(ref reflect.Value) error {
	refType := ref.Type()
	var errorList []string
	
	for i := 0; i < refType.NumField(); i++ {
		refField := ref.Field(i)
		if reflect.Ptr == refField.Kind() && !refField.IsNil() && refField.CanSet() {
			err := parse(refField.Interface())
			if nil != err {
				return err
			}
			continue
		}
		refTypeField := refType.Field(i)
		value := refTypeField.Tag.Get("env")
		if value == "" {
			continue
		}
		if err := set(refField, refTypeField, value); err != nil {
			errorList = append(errorList, err.Error())
			continue
		}
	}
	
	if len(errorList) == 0 {
		return nil
	}
	return errors.New(strings.Join(errorList, ","))
}

func set(field reflect.Value, refType reflect.StructField, value string) error {
	switch field.Kind() {
	case reflect.String:
		field.SetString(viper.GetString(value))
	case reflect.Bool:
		field.SetBool(viper.GetBool(value))
	case reflect.Int:
		field.SetInt(int64(viper.GetInt(value)))
	default:
		return handleTextUnmarshaler(field, value)
	}
	return nil
}

func handleTextUnmarshaler(field reflect.Value, value string) error {
	if reflect.Ptr == field.Kind() {
		if field.IsNil() {
			field.Set(reflect.New(field.Type().Elem()))
		}
	} else if field.CanAddr() {
		field = field.Addr()
	}
	
	tm, ok := field.Interface().(encoding.TextUnmarshaler)
	if !ok {
		return ErrUnsupportedType
	}
	
	return tm.UnmarshalText([]byte(value))
}
