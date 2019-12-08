package main

type BranchRange struct {
	Min int
	Max int
}

type Bank struct {
	Name string
	Code int
	Branches []BranchRange
}
func (b *Bank) FirstBranch() BranchRange {
	return b.Branches[0]
}


func Banks() map[int]Bank {
	return map[int]Bank{
		1: Bank{
			Name:"ANZ", 
			Code: 1, 
			Branches: []BranchRange{
						BranchRange{Min: 1, Max: 5699},
					},
		},		
		2: Bank{
				Name:"BNZ", 
				Code: 2, 
				Branches: []BranchRange{
							BranchRange{Min: 1, Max: 999},
							BranchRange{Min: 1200, Max: 1299},
						},
			},
		3: Bank{
			Name:"Westpac", 
			Code: 3, 
			Branches: []BranchRange{
						BranchRange{Min: 1, Max: 1999},
					},
		},			
		6: Bank{
			Name:"ANZ", 
			Code: 3, 
			Branches: []BranchRange{
						BranchRange{Min: 1, Max: 1499},
					},
		},				
		8: Bank{
			Name:"NAB", 
			Code: 8, 
			Branches: []BranchRange{
						BranchRange{Min: 0, Max: 9999},
					},
		},
		10: Bank{
			Name:"Industrial and Commerical Bank of China", 
			Code: 10, 
			Branches: []BranchRange{
						BranchRange{Min: 5165, Max: 5169},
					},
		},		
		11: Bank{
			Name:"Postbank", 
			Code: 11, 
			Branches: []BranchRange{
						BranchRange{Min: 5000, Max: 8999},
					},
		},			
		12: Bank{
			Name:"ASB", 
			Code: 12, 
			Branches: []BranchRange{
						BranchRange{Min: 3000, Max: 3999},
					},
		},			
		13: Bank{
			Name:"Trust Bank Southland", 
			Code: 14, 
			Branches: []BranchRange{
						BranchRange{Min: 4900, Max: 4999},
					},
		},		
		14: Bank{
			Name:"Trust Bank Otago", 
			Code: 14, 
			Branches: []BranchRange{
						BranchRange{Min: 4700, Max: 4799},
					},
		},		
		15: Bank{
			Name:"TSB", 
			Code: 15, 
			Branches: []BranchRange{
						BranchRange{Min: 3900, Max: 3999},
					},
		},	
		16: Bank{
			Name:"Trust Bank Cantebury", 
			Code: 16, 
			Branches: []BranchRange{
						BranchRange{Min: 4400, Max: 4499},
					},
		},		
		17: Bank{
			Name:"Trust Bank Waikato", 
			Code: 17, 
			Branches: []BranchRange{
						BranchRange{Min: 3300, Max: 3399},
					},
		},	
		18: Bank{
			Name:"Trust Bank Bay of Plenty", 
			Code: 18, 
			Branches: []BranchRange{
						BranchRange{Min: 3500, Max: 3599},
					},
		},		
		19: Bank{
			Name:"Trust Bank South Cantebury", 
			Code: 19, 
			Branches: []BranchRange{
						BranchRange{Min: 4600, Max: 4699},
					},
		},
		20: Bank{
			Name:"Trust Bank Central", 
			Code: 20, 
			Branches: []BranchRange{
						BranchRange{Min: 4100, Max: 4199},
					},
		},		
		21: Bank{
			Name:"Trust Bank Auckland", 
			Code: 21, 
			Branches: []BranchRange{
						BranchRange{Min: 4800, Max: 4899},
					},
		},	
		22: Bank{
			Name:"Trust Bank Wanganui", 
			Code: 22, 
			Branches: []BranchRange{
						BranchRange{Min: 4000, Max: 4049},
					},
		},			
		23: Bank{
			Name:"Trust Bank Wellington", 
			Code: 23, 
			Branches: []BranchRange{
						BranchRange{Min: 3700, Max: 3799},
					},
		},		
		24: Bank{
			Name:"Westland", 
			Code: 24, 
			Branches: []BranchRange{
						BranchRange{Min: 4300, Max: 4349},
					},
		},	
		25: Bank{
			Name:"Countrywide", 
			Code: 25, 
			Branches: []BranchRange{
						BranchRange{Min: 2500, Max: 2599},
					},
		},			
		30: Bank{
			Name:"HSBC", 
			Code: 30, 
			Branches: []BranchRange{
						BranchRange{Min: 2900, Max: 2956},
					},
		},			
		31: Bank{
			Name:"Citibank", 
			Code: 31, 
			Branches: []BranchRange{
						BranchRange{Min: 2900, Max: 2956},
					},
		},	
		38: Bank{
			Name:"Kiwibank", 
			Code: 38, 
			Branches: []BranchRange{
						BranchRange{Min: 9000, Max: 9499},
					},
		},														
	}
}

func CheckDigitWeighting(bank int, account int) []int {
	switch (bank) {
		case 8:
			return []int{0,0,0,0,0,0,0,7,6,5,4,3,2,1,0,0,0,0}
		case 9:
			return []int{0,0,0,0,0,0,0,0,0,0,5,4,3,2,0,0,0,1}
		case 25:
		case 33:
			return []int{0,0,0,0,0,0,0,1,7,3,1,7,3,1,0,0,0,0}
		case 26:
		case 28:
		case 29:
			return []int{0,0,0,0,0,0,0,1,3,7,1,3,7,1,0,3,7,1}
		case 31:
			return []int{}
  	}	
	
	if account < 990000 {
		return []int{0,0,6,3,7,9,0,0,10,5,8,4,2,1,0,0,0,0}
	}

	return []int{0,0,0,0,0,0,0,0,10,5,8,4,2,1,0,0,0,0}  
}

func Modulo(bank int) int {
	switch (bank) {
		case 25:
		case 33:
		case 26:
		case 28:
		case 29:
			return 10
		case 31:
			return 1
  }

  return 11
}