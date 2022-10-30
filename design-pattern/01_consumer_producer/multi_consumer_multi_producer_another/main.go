package main

import (
	"fmt"
	"sync"
)

var quotes = [][]string{
	{
		"你要尽全力保护你的梦想。那些嘲笑你的人，他们必定会失败，他们想把你变成和他们一样的人。如果你有梦想的话，就要努力去实现。",
		"如果做人没有梦想，那跟咸鱼有什么分别？",
		"死亡不是生命的终点，遗忘才是。",
		"星星在哪里都是很亮的，就看你有没有抬头去看他们。",
	},
	{
		"不管前方的路有多苦，只要走的方向正确，不管多么崎岖不平，都比站在原地更加接近幸福。",
		"只要你肯领略，就会发现人生本是多么可爱，每个季节里有很多足以让你忘记所有烦恼的赏心乐趣。",
		"开拓视野，冲破艰险，看见世界，身临其境，贴近彼此，感受生活，这就是生活的目的。",
		"我们读诗写诗，非为它的灵巧。我们读诗写诗，因为我们是人类的一员。而人类充满了热情。医药，法律，商业，工程，这些都是高贵的理想，并且是维生的必需条件。但是诗，美，浪漫，爱，这些才是我们生存的原因。",
	},
	{
		"如果你不出去走走，你就会以为这就是全世界。",
		"如果我不顾一切发挥每一点潜能去做会怎样？我必须做到，我别无选择。",
		"你真正是谁并不重要，重要的是你的所做所为。",
		"有信心不一定会成功，没信心一定不会成功。",
	},
	{
		"决定我们成为什么样人的，不是我们的能力，而是我们的选择。",
		"有时候你只需要花二十秒，疯狂地一鼓作气。仅仅花上二十秒，鼓起勇气，即便有多尴尬。然后我向你保证，会有好事发生的。",
	},
}

const producerCount int = 4
const consumerCount int = 3

func produce(link chan<- string, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, msg := range quotes[id] {
		link <- msg
	}
}

func consume(link <-chan string, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range link {
		fmt.Printf("Message \"%v\" is consumed by consumer %v\n", msg, id)
	}
}

func main() {
	link := make(chan string)
	wp := &sync.WaitGroup{}
	wc := &sync.WaitGroup{}

	wp.Add(producerCount)
	wc.Add(consumerCount)

	for i := 0; i < producerCount; i++ {
		go produce(link, i, wp)
	}

	for i := 0; i < consumerCount; i++ {
		go consume(link, i, wc)
	}

	wp.Wait()
	close(link)
	wc.Wait()
}
