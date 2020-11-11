package main

func main(){
	times := runCases(10,"usbrews.json")
	//Produces 5 cases with response = [0,..,8] and returns an array containing the average time for each after n runs per case

	for _,avgTime:= range times{
		println(avgTime)
	}
}