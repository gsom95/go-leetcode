package reversewordsinastring3

import (
	"math/rand"
	"strings"
	"testing"
)

func Test_reverseWords(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "example 1",
			s:    "Let's take LeetCode contest",
			want: "s'teL ekat edoCteeL tsetnoc",
		},
		{
			name: "example 2",
			s:    "God Ding",
			want: "doG gniD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.s); got != tt.want {
				t.Errorf("reverseWords() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

var testStrings = func() []string {
	return strings.Split(`Lorem ipsum dolor sit amet consectetur adipiscing elit. Nunc elementum est quis nunc finibus tempus. Aliquam condimentum sapien fermentum urna ullamcorper eleifend molestie elit tempor. Pellentesque molestie ultricies quam quis congue. Nullam leo turpis hendrerit ac turpis nec sodales auctor nunc. Morbi fermentum ipsum ut magna sodales et pulvinar arcu varius. Sed ac auctor diam. Phasellus ut nisl pulvinar lobortis risus quis cursus eros. Nunc mollis sit amet felis vel viverra. Morbi vel maximus mi. Phasellus iaculis ipsum eu mauris luctus lacinia. Quisque varius elit aliquam odio tincidunt sagittis euismod sit amet neque. Aenean congue mauris non erat tempor ultricies. Donec bibendum suscipit imperdiet. Nunc sed metus faucibus condimentum nibh non accumsan purus. Cras rhoncus neque quis porta lacinia. Nam mattis nisl velit non eleifend nisl imperdiet vitae. Praesent aliquam eros sit amet ligula sollicitudin eu lacinia nulla sagittis. Aliquam non interdum ante vitae blandit lacus. Etiam convallis sapien non hendrerit semper elit nisi lacinia quam ut dignissim augue leo ac nibh. Nullam lorem nulla fringilla ac rhoncus id consectetur ac felis. Nam porttitor egestas dolor id maximus. Curabitur scelerisque augue non ornare ultrices ex mi iaculis ante quis rutrum magna felis nec mauris. Vivamus rutrum magna eget pharetra tempus. Praesent pretium magna in tempor congue. In semper ex ac nisl pulvinar nec efficitur odio ornare. Vivamus ipsum justo rhoncus vitae enim eu tincidunt lacinia augue. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Suspendisse et velit eu tellus blandit iaculis. Aliquam convallis ligula sapien eget imperdiet lacus auctor nec. Suspendisse pretium massa dolor at porta ex iaculis ut. Proin volutpat elit at sodales lacinia massa tellus volutpat orci eget dictum nibh nulla a risus. Vestibulum dictum accumsan eros tincidunt venenatis lacus tincidunt laoreet. Aenean ultrices odio quis nisl faucibus consequat. Sed eu urna sed nisl convallis molestie sed ut orci. Aliquam venenatis suscipit purus vitae dictum. Suspendisse vitae finibus ligula pulvinar pretium tortor. Sed ornare finibus mi ac vehicula. Ut lorem mi feugiat fermentum ligula at bibendum efficitur ipsum. Donec ac ornare arcu eu laoreet felis. Sed tristique neque non blandit efficitur. Aenean maximus facilisis sapien nec lobortis. Etiam non porttitor leo. Phasellus et elit id tortor tincidunt vestibulum. Donec ipsum lacus ultrices sed lectus et egestas hendrerit lectus. Vivamus scelerisque arcu in diam consectetur malesuada. Duis congue dui iaculis elementum lorem vel rhoncus velit. Etiam et porta turpis. Suspendisse porta nec dolor et feugiat. Phasellus sagittis tincidunt molestie. Sed a velit vel sem consectetur pretium. Duis quis dolor eget felis tincidunt molestie eget et arcu.`, ".")
}()

var global string

func BenchmarkReverseWithNewBulder(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		testString := testStrings[rand.Intn(len(testStrings))]
		s = reverseWordsWithNewBulder(testString)
	}
	global = s
}

func BenchmarkReverseWithBulderReset(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		testString := testStrings[rand.Intn(len(testStrings))]
		s = reverseWordsWithReset(testString)
	}
	global = s
}
