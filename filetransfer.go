package main
import( 
	"os"
"io"
"log"
)
func main() {
   sourceLocation, err := os.Open ("C:/Users/kota/Documents/hello.docx")
     if err != nil{
		 log.Fatal(err)
	 }
	 defer sourceLocation.Close()
   destinationLocation, err := os.Create("E:/test/hello.docx")
   if err != nil{
	   log.Panic(err)
   }
   defer destinationLocation.Close()
   finalLocation, err := io.Copy(destinationLocation,sourceLocation)
    if err != nil {
		log.Println(err)
		
	}
   log.Println(finalLocation)
}
