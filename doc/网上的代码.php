<?php 
//设置最大执行时间
set_time_limit(0);
function getHtml($url){
  // 1. 初始化
   $ch = curl_init();
   // 2. 设置选项，包括URL
   curl_setopt($ch,CURLOPT_URL,$url);
   curl_setopt($ch,CURLOPT_RETURNTRANSFER,1);
   curl_setopt($ch,CURLOPT_HEADER,0);
   // 3. 执行并获取HTML文档内容
   $output = curl_exec($ch);
   if($output === FALSE ){
    $output = '';
   }
   // 4. 释放curl句柄
   curl_close($ch);
   return $output;
}
function getPageData($url){
  // 获取整个网页内容
  $html = getHtml($url);
  // 初步获取主块内容
  preg_match("/教程列表.*教程列表/s",$html,$body_html);
  // 返回数据
  $data = array();
  //判断是否存在要获取的内容
  if(count($body_html)){
    // 获取页面指定信息
    preg_match_all('/<a class="avatar".*user_id="(\S*)" href="(\S*)" rel="external nofollow" /',$body_html[0],$info_1);
    preg_match_all('/<a href="(.*)" rel="external nofollow" .*title="(.*)"/',$body_html[0],$info_2);
    $info = array_merge($info_1,$info_2);
    //组合的信息
    for($index=0; $index<count($info[0]); $index++){
      //以文章信息作为key存数组，以及覆盖旧数据
      $data[$info[4][$index]] = array(
              'user_id'  => $info[1][$index],
              'user_home' => $info[2][$index],
              'a_url'   => $info[4][$index],
              'a_title'  => $info[5][$index],
           );
    }
  }
  return $data;
}
header("Content-type: text/html; charset=utf-8"); 
echo '<pre>';
// 初始化数据
$page_no = 1;
$data_all = array();
// 分页获取数据
do{
  $url = 'http://www.thinkphp.cn/code/examples/p/' . $page_no;
  $data = getPageData($url);
  $data_all += $data;
  $page_no ++;
}while ($page_no <= 10); //当前只获取10页，如果要全部获取则把条件换成$data或!empty($data)
var_dump($data_all);
?>