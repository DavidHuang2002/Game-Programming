package main
import ("fmt"
"bufio"
"os"
"strconv")

type piece int

const (
    blank piece = iota
    x
    o
    cat
)








type TicTacToeBoard struct {
	board []piece
}

func NewTicTacToeBoard() *TicTacToeBoard {
    return &TicTacToeBoard{[]piece{blank,blank,blank,blank,blank,blank,blank,blank,blank}}
}

func (b TicTacToeBoard) String() string {
    s:=""
    for i,v :=  range b.board {
        if v==blank {
            s+="   "
        } else if v==x {
            s+=" x "
        } else if v==o {
            s+=" o "
        }
        
        if i%3<2 {
            s+="|"
        }
        
        if i%3==2 && i/3<2 {
            s+="\n---+---+---\n"
        }
    }
    return s//fmt.Sprintf("%", b)
}

func (b *TicTacToeBoard) place(location int){
    b.board[location]=b.getPlayer()
}

func (b *TicTacToeBoard) getPlayer() piece {
    c:=9
    for _,v :=  range b.board {
        if v==blank {
            c--
        }
    }
    if c==9{
		return cat
	}
    if c%2==0 {
        return x
    }
    return o
}

func (b *TicTacToeBoard) getPieceLoc(piese piece) []int{
	var temp []int
	for i,v := range b.board{
		if v==piese{
			temp=append(temp,i)
		}
	}
	return temp
}

func isIn(a int, b []int) bool{
	for _,v :=range b{
		if a==v {
			return true
		}
	}
	return false
}

func  arraryIsIn(a []int, b []int) bool{
	flag:=true
	for _,v := range a{
		flag=isIn(v,b)
		if flag!=true{
			return false
		}
	}
	return flag
}
func winCheck(){
	
}

func (b *TicTacToeBoard) getWinner() piece {
	xLoc:=b.getPieceLoc(x)
	oLoc:=b.getPieceLoc(o)
	winningCondition:=[][]int{[]int{0,3,6},
		[]int{1,4,7},
		[]int{2,5,8},
		[]int{0,1,2},
		[]int{3,4,5},
		[]int{6,7,8},
		[]int{0,4,8},
		[]int{2,4,6}}
    for _,v:=range winningCondition{
		if arraryIsIn(v,xLoc){
			return x
		}
		if arraryIsIn(v,oLoc){
			return o
		}
	}
    return blank
}



func main() {
	reader:=bufio.NewReader(os.Stdin)
	outerer:
	for 1>0{
	board:=NewTicTacToeBoard()
	outer: 
	for 1>0{
		current_player:=board.getPlayer()
		if current_player==cat{
			fmt.Println("Draw, you losers.")
			fmt.Println(board)
			break
		}
		fmt.Println(board)
		fmt.Print("It is player ")
		fmt.Print(current_player)
		fmt.Println("'s turn")
		fmt.Println("Please input where you want to place your piece")
		input,_:=reader.ReadString('\n')
		loc,_:=strconv.Atoi(input[:1])
		board.place(loc)
		switch board.getWinner(){
			case 1:
				fmt.Println("X: Player 1 win")
				fmt.Println(board)
				break outer
			case 2:
				fmt.Println("O: Player 2 win")
				fmt.Println(board)
				break outer
		}
	}
	for true{
	fmt.Println("Input Y to keep playing, Q to quit.")
	input,_:=reader.ReadString('\n')
	if input == "Y\n"{
		continue outerer
	}else if input =="Q\n"{
		break outerer
	}else{
		fmt.Println("Not a valid input")
	}
	}
	}
}
