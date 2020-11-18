package oauth

/*
// uncomment to test the server

func TestStartServer(t *testing.T) {
	d := new(AppData)
	d.Name = "Testing123"
	com := make(chan interface{})
	StartServer(d, "localhost:8080", com)
OUT:
	for {
		what := <-com // blocks
		switch v := what.(type) {
		case string:
			switch v {
			case "killmenow":
				t.Log("Killing server.")
				break OUT
			}
		}
	}
}
*/
