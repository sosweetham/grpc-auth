package global

import (
	"os"

	"github.com/sohamjaiswal/grpc-ftp/tools"
)

var DefaultHasher = &tools.Hasher{
    Salt:        os.Getenv("PASS_SALT"),  // You should generate a secure salt
	Iterations:  1,
    Memory:      65536,
    Parallelism: 4,
    KeyLength:   32,
}