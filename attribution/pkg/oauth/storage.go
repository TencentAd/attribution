package oauth

var (
    Prefix = "access_token"
)

type store interface {
    Get(string) (string, error)
    Set(string, string) error
}
