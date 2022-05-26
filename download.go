import (
	"fmt"
	"os"
	"log"
	"github.com/cavaliercoder/grab"
)
func main() {
	resp, err := grab.Get(".", "http://www.golang-book.com/public/pdf/gobook.pdf")
	if err != nil {
		log.Fatal(err)
	}
}
fmt.Println("Download saved to", resp.Filename)