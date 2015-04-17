package main
import (
	"subHtml"
	"fmt"
)

func main() {
	htmlStr := `<p style="margin-top: 15px; margin-bottom: 0px; padding: 0px; color: rgb(68, 68, 68); font-family: &#39;Helvetica Neue&#39;, Helvetica, Arial, sans-serif; font-size: 14px; line-height: 22px; text-align: justify; white-space: normal;">在项目中遇到了以下需求：golang的webserver要将页面中用户提交的文件等信息做中转，提交到云存储服务器中。</p><p style="margin-top: 15px; margin-bottom: 0px; padding: 0px; color: rgb(68, 68, 68); font-family: &#39;Helvetica Neue&#39;, Helvetica, Arial, sans-serif; font-size: 14px; line-height: 22px; text-align: justify; white-space: normal;">1、折中方案：先将用户提交的文件等信息存储到服务器本地，再使用beego中httplib.PostFile将本地的文件上传到服务器。</p><p style="margin-top: 15px; margin-bottom: 0px; padding: 0px; color: rgb(68, 68, 68); font-family: &#39;Helvetica Neue&#39;, Helvetica, Arial, sans-serif; font-size: 14px; line-height: 22px; text-align: justify; white-space: normal;">2、最终方案：服务端拼装form数据，做数据中转，模拟提交。<br style="margin: 0px; padding: 0px;"/><a id="more" style="margin: 0px; padding: 0px; color: rgb(37, 143, 184);"></a></p><p style="margin-top: 15px; margin-bottom: 0px; padding: 0px; color: rgb(68, 68, 68); font-family: &#39;Helvetica Neue&#39;, Helvetica, Arial, sans-serif; font-size: 14px; line-height: 22px; text-align: justify; white-space: normal;">方案1、先看折中方案的实施：<br style="margin: 0px; padding: 0px;"/> <br style="margin: 0px; padding: 0px;"/>part1:文件存储到服务器【略】</p><p style="margin-top: 15px; margin-bottom: 0px; padding: 0px; color: rgb(68, 68, 68); font-family: &#39;Helvetica Neue&#39;, Helvetica, Arial, sans-serif; font-size: 14px; line-height: 22px; text-align: justify; white-space: normal;">part2:文件中转</p><pre class="brush:go;toolbar:false">req:=httplib.POST(&quot;url&quot;)
req.Header(&quot;myheader&quot;,&quot;header&quot;)<p style="margin-t
	`

	textStr := `在项目中遇到了以下需求：golang的webserver要将页面中用户提交的文件等信息做中转，提交到云存储服务器中。
1、折中方案：先将用户提交的文件等信息存储到服务器本地，再使用beego中httplib.PostFile将本地的文件上传到服务器。
2、最终方案：服务端拼装form数据，做数据中转，模拟提交。
方案1、先看折中方案的实施：
part1:文件存储到服务器【略】
part2:文件中转
`

	fmt.Println("before1", htmlStr)
	fmt.Println("after1", subHtml.SubHtml(htmlStr, 50)+"\n")

	fmt.Println("before2", textStr)
	fmt.Println("after2", subHtml.SubHtml(textStr, 50))
}


