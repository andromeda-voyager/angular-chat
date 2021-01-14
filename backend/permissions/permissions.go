package permissions

const (
	Full         = 255
	None         = 0
	DeleteServer = 128
)

// i := int64(255)
// k := ^(i & 0)
// fmt.Println(k)
// fmt.Printf("%b", i)

func CanDeleteServer(permissions uint8) bool {
	return (permissions & DeleteServer) == DeleteServer
}
