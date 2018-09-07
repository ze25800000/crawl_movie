package controllers

import (
	"github.com/astaxie/beego"
	"crawl_movie/models"
)

type CrawlMovieController struct {
	beego.Controller
}

func (c *CrawlMovieController) CrawlMovie() {
	sMovieHtml := `
<!DOCTYPE html>
<html lang="zh-cmn-Hans" class="ua-windows ua-webkit">
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="renderer" content="webkit">
    <meta name="referrer" content="always">
    <meta name="google-site-verification" content="ok0wCgT20tBBgo9_zat2iAcimtN4Ftf5ccsh092Xeyw" />
    <title>
        碟中谍6：全面瓦解 (豆瓣)
</title>
    
    <meta name="baidu-site-verification" content="cZdR4xxR7RxmM4zE" />
    <meta http-equiv="Pragma" content="no-cache">
    <meta http-equiv="Expires" content="Sun, 6 Mar 2005 01:00:00 GMT">
    
    <link rel="apple-touch-icon" href="https://img3.doubanio.com/f/movie/d59b2715fdea4968a450ee5f6c95c7d7a2030065/pics/movie/apple-touch-icon.png">
    <link href="https://img3.doubanio.com/f/shire/bf61b1fa02f564a4a8f809da7c7179b883a56146/css/douban.css" rel="stylesheet" type="text/css">
    <link href="https://img3.doubanio.com/f/shire/ae3f5a3e3085968370b1fc63afcecb22d3284848/css/separation/_all.css" rel="stylesheet" type="text/css">
    <link href="https://img3.doubanio.com/f/movie/8864d3756094f5272d3c93e30ee2e324665855b0/css/movie/base/init.css" rel="stylesheet">
    <script type="text/javascript">var _head_start = new Date();</script>
    <script type="text/javascript" src="https://img3.doubanio.com/f/movie/0495cb173e298c28593766009c7b0a953246c5b5/js/movie/lib/jquery.js"></script>
    <script type="text/javascript" src="https://img3.doubanio.com/f/shire/f010949d3f23dd7c972ad7cb40b800bf70723c93/js/douban.js"></script>
    <script type="text/javascript" src="https://img3.doubanio.com/f/shire/0efdc63b77f895eaf85281fb0e44d435c6239a3f/js/separation/_all.js"></script>
    
    <meta name="keywords" content="碟中谍6：全面瓦解,Mission: Impossible - Fallout,碟中谍6：全面瓦解影评,剧情介绍,电影图片,预告片,影讯,在线购票,论坛">
    <meta name="description" content="碟中谍6：全面瓦解电影简介和剧情介绍,碟中谍6：全面瓦解影评、图片、预告片、影讯、论坛、在线购票">
    <meta name="mobile-agent" content="format=html5; url=http://m.douban.com/movie/subject/26336252/"/>
    <link rel="alternate" href="android-app://com.douban.movie/doubanmovie/subject/26336252/" />
    <link rel="stylesheet" href="https://img3.doubanio.com/dae/cdnlib/libs/LikeButton/1.0.5/style.min.css">
    <script type="text/javascript" src="https://img3.doubanio.com/f/shire/77323ae72a612bba8b65f845491513ff3329b1bb/js/do.js" data-cfg-autoload="false"></script>
    <script type="text/javascript">
      Do.add('dialog', {path: 'https://img3.doubanio.com/f/shire/4ea3216519a6183c7bcd4f7d1a6d4fd57ce1a244/js/ui/dialog.js', type: 'js'});
      Do.add('dialog-css', {path: 'https://img3.doubanio.com/f/shire/8377b9498330a2e6f056d863987cc7a37eb4d486/css/ui/dialog.css', type: 'css'});
      Do.add('handlebarsjs', {path: 'https://img3.doubanio.com/f/movie/3d4f8e4a8918718256450eb6e57ec8e1f7a2e14b/js/movie/lib/handlebars.current.js', type: 'js'});
    </script>
    
  <script type='text/javascript'>
  var _vwo_code = (function() {
    var account_id = 249272,
      settings_tolerance = 0,
      library_tolerance = 2500,
      use_existing_jquery = false,
      // DO NOT EDIT BELOW THIS LINE
      f=false,d=document;return{use_existing_jquery:function(){return use_existing_jquery;},library_tolerance:function(){return library_tolerance;},finish:function(){if(!f){f=true;var a=d.getElementById('_vis_opt_path_hides');if(a)a.parentNode.removeChild(a);}},finished:function(){return f;},load:function(a){var b=d.createElement('script');b.src=a;b.type='text/javascript';b.innerText;b.onerror=function(){_vwo_code.finish();};d.getElementsByTagName('head')[0].appendChild(b);},init:function(){settings_timer=setTimeout('_vwo_code.finish()',settings_tolerance);var a=d.createElement('style'),b='body{opacity:0 !important;filter:alpha(opacity=0) !important;background:none !important;}',h=d.getElementsByTagName('head')[0];a.setAttribute('id','_vis_opt_path_hides');a.setAttribute('type','text/css');if(a.styleSheet)a.styleSheet.cssText=b;else a.appendChild(d.createTextNode(b));h.appendChild(a);this.load('//dev.visualwebsiteoptimizer.com/j.php?a='+account_id+'&u='+encodeURIComponent(d.URL)+'&r='+Math.random());return settings_timer;}};}());

  +function () {
    var bindEvent = function (el, type, handler) {
        var $ = window.jQuery || window.Zepto || window.$
       if ($ && $.fn && $.fn.on) {
           $(el).on(type, handler)
       } else if($ && $.fn && $.fn.bind) {
           $(el).bind(type, handler)
       } else if (el.addEventListener){
         el.addEventListener(type, handler, false);
       } else if (el.attachEvent){
         el.attachEvent("on" + type, handler);
       } else {
         el["on" + type] = handler;
       }
     }

    var _origin_load = _vwo_code.load
    _vwo_code.load = function () {
      var args = [].slice.call(arguments)
      bindEvent(window, 'load', function () {
        _origin_load.apply(_vwo_code, args)
      })
    }
  }()

  _vwo_settings_timer = _vwo_code.init();
  </script>


    <style type="text/css">
  
</style>
    <style type="text/css">img { max-width: 100%; }</style>
    <script type="text/javascript"></script>
    <link rel="stylesheet" href="https://img3.doubanio.com/misc/mixed_static/e421d55b2b427d7.css">

    <link rel="shortcut icon" href="https://img3.doubanio.com/favicon.ico" type="image/x-icon">
</head>

<body>
  
    <script type="text/javascript">var _body_start = new Date();</script>

    
    



    <link href="//img3.doubanio.com/dae/accounts/resources/984c231/shire/bundle.css" rel="stylesheet" type="text/css">



<div id="db-global-nav" class="global-nav">
  <div class="bd">
    
<div class="top-nav-info">
  <a href="https://www.douban.com/accounts/login?source=movie" class="nav-login" rel="nofollow">登录</a>
  <a href="https://www.douban.com/accounts/register?source=movie" class="nav-register" rel="nofollow">注册</a>
</div>


    <div class="top-nav-doubanapp">
  <a href="https://www.douban.com/doubanapp/app?channel=top-nav" class="lnk-doubanapp">下载豆瓣客户端</a>
  <div id="doubanapp-tip">
    <a href="https://www.douban.com/doubanapp/app?channel=qipao" class="tip-link">豆瓣 <span class="version">6.0</span> 全新发布</a>
    <a href="javascript: void 0;" class="tip-close">×</a>
  </div>
  <div id="top-nav-appintro" class="more-items">
    <p class="appintro-title">豆瓣</p>
    <p class="qrcode">扫码直接下载</p>
    <div class="download">
      <a href="https://www.douban.com/doubanapp/redirect?channel=top-nav&direct_dl=1&download=iOS">iPhone</a>
      <span>·</span>
      <a href="https://www.douban.com/doubanapp/redirect?channel=top-nav&direct_dl=1&download=Android" class="download-android">Android</a>
    </div>
  </div>
</div>

    


<div class="global-nav-items">
  <ul>
    <li class="">
      <a href="https://www.douban.com" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-main&quot;,&quot;uid&quot;:&quot;0&quot;}">豆瓣</a>
    </li>
    <li class="">
      <a href="https://book.douban.com" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-book&quot;,&quot;uid&quot;:&quot;0&quot;}">读书</a>
    </li>
    <li class="on">
      <a href="https://movie.douban.com"  data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-movie&quot;,&quot;uid&quot;:&quot;0&quot;}">电影</a>
    </li>
    <li class="">
      <a href="https://music.douban.com" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-music&quot;,&quot;uid&quot;:&quot;0&quot;}">音乐</a>
    </li>
    <li class="">
      <a href="https://www.douban.com/location" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-location&quot;,&quot;uid&quot;:&quot;0&quot;}">同城</a>
    </li>
    <li class="">
      <a href="https://www.douban.com/group" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-group&quot;,&quot;uid&quot;:&quot;0&quot;}">小组</a>
    </li>
    <li class="">
      <a href="https://read.douban.com&#47;?dcs=top-nav&amp;dcm=douban" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-read&quot;,&quot;uid&quot;:&quot;0&quot;}">阅读</a>
    </li>
    <li class="">
      <a href="https://douban.fm&#47;?from_=shire_top_nav" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-fm&quot;,&quot;uid&quot;:&quot;0&quot;}">FM</a>
    </li>
    <li class="">
      <a href="https://time.douban.com&#47;?dt_time_source=douban-web_top_nav" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-time&quot;,&quot;uid&quot;:&quot;0&quot;}">时间</a>
    </li>
    <li class="">
      <a href="https://market.douban.com&#47;?utm_campaign=douban_top_nav&amp;utm_source=douban&amp;utm_medium=pc_web" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-market&quot;,&quot;uid&quot;:&quot;0&quot;}">豆品</a>
    </li>
    <li>
      <a href="#more" class="bn-more"><span>更多</span></a>
      <div class="more-items">
        <table cellpadding="0" cellspacing="0">
          <tbody>
            <tr>
              <td>
                <a href="https://ypy.douban.com" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-ypy&quot;,&quot;uid&quot;:&quot;0&quot;}">豆瓣摄影</a>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </li>
  </ul>
</div>

  </div>
</div>
<script>
  ;window._GLOBAL_NAV = {
    DOUBAN_URL: "https://www.douban.com",
    N_NEW_NOTIS: 0,
    N_NEW_DOUMAIL: 0
  };
</script>



    <script src="//img3.doubanio.com/dae/accounts/resources/984c231/shire/bundle.js" defer="defer"></script>




    



    <link href="//img3.doubanio.com/dae/accounts/resources/19507ad/movie/bundle.css" rel="stylesheet" type="text/css">




<div id="db-nav-movie" class="nav">
  <div class="nav-wrap">
  <div class="nav-primary">
    <div class="nav-logo">
      <a href="https:&#47;&#47;movie.douban.com">豆瓣电影</a>
    </div>
    <div class="nav-search">
      <form action="https:&#47;&#47;movie.douban.com/subject_search" method="get">
        <fieldset>
          <legend>搜索：</legend>
          <label for="inp-query">
          </label>
          <div class="inp"><input id="inp-query" name="search_text" size="22" maxlength="60" placeholder="搜索电影、电视剧、综艺、影人" value=""></div>
          <div class="inp-btn"><input type="submit" value="搜索"></div>
          <input type="hidden" name="cat" value="1002" />
        </fieldset>
      </form>
    </div>
  </div>
  </div>
  <div class="nav-secondary">
    

<div class="nav-items">
  <ul>
    <li    ><a href="https://movie.douban.com/cinema/nowplaying/"
     >影讯&购票</a>
    </li>
    <li    ><a href="https://movie.douban.com/explore"
     >选电影</a>
    </li>
    <li    ><a href="https://movie.douban.com/tv/"
     >电视剧</a>
    </li>
    <li    ><a href="https://movie.douban.com/chart"
     >排行榜</a>
    </li>
    <li    ><a href="https://movie.douban.com/tag/"
     >分类</a>
    </li>
    <li    ><a href="https://movie.douban.com/review/best/"
     >影评</a>
    </li>
    <li    ><a href="https://movie.douban.com/annual/2017?source=navigation"
     >2017年度榜单</a>
    </li>
    <li    ><a href="https://movie.douban.com/standbyme/2017?source=navigation"
     >2017观影报告</a>
    </li>
  </ul>
</div>

  </div>
</div>

<script id="suggResult" type="text/x-jquery-tmpl">
  <li data-link="{{= url}}">
            <a href="{{= url}}" onclick="moreurl(this, {from:'movie_search_sugg', query:'{{= keyword }}', subject_id:'{{= id}}', i: '{{= index}}', type: '{{= type}}'})">
            <img src="{{= img}}" width="40" />
            <p>
                <em>{{= title}}</em>
                {{if year}}
                    <span>{{= year}}</span>
                {{/if}}
                {{if sub_title}}
                    <br /><span>{{= sub_title}}</span>
                {{/if}}
                {{if address}}
                    <br /><span>{{= address}}</span>
                {{/if}}
                {{if episode}}
                    {{if episode=="unknow"}}
                        <br /><span>集数未知</span>
                    {{else}}
                        <br /><span>共{{= episode}}集</span>
                    {{/if}}
                {{/if}}
            </p>
        </a>
        </li>
  </script>




    <script src="//img3.doubanio.com/dae/accounts/resources/19507ad/movie/bundle.js" defer="defer"></script>





    
    <div id="wrapper">
        

        
    <div id="content">
        

    <div id="dale_movie_subject_top_icon"></div>
    <h1>
        <span property="v:itemreviewed">碟中谍6：全面瓦解 Mission: Impossible - Fallout</span>
            <span class="year">(2018)</span>
    </h1>

        <div class="grid-16-8 clearfix">
            

            
            <div class="article">
                
    

    





        <div class="indent clearfix">
            <div class="subjectwrap clearfix" xmlns:v="http://rdf.data-vocabulary.org/#" typeof="v:Movie">
                <div class="subject clearfix">
                    



<div id="mainpic" class="">
    <a class="nbgnbg" href="https://movie.douban.com/subject/26336252/photos?type=R" title="点击看更多海报">
        <img src="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2529365085.webp" title="点击看更多海报" alt="Mission: Impossible - Fallout" rel="v:image" />
   </a>
</div>

                    


<div id="info">
        <span ><span class='pl'>导演</span>: <span class='attrs'><a href="/celebrity/1276314/" rel="v:directedBy">克里斯托弗·麦奎里</a></span></span><br/>
        <span ><span class='pl'>编剧</span>: <span class='attrs'><a href="/celebrity/1276314/">克里斯托弗·麦奎里</a> / <a href="/celebrity/1028370/">布鲁斯·盖勒</a></span></span><br/>
        <span class="actor"><span class='pl'>主演</span>: <span class='attrs'><a href="/celebrity/1054435/" rel="v:starring">汤姆·克鲁斯</a> / <a href="/celebrity/1044713/" rel="v:starring">亨利·卡维尔</a> / <a href="/celebrity/1048129/" rel="v:starring">文·瑞姆斯</a> / <a href="/celebrity/1035648/" rel="v:starring">西蒙·佩吉</a> / <a href="/celebrity/1088314/" rel="v:starring">丽贝卡·弗格森</a> / <a href="/celebrity/1286119/" rel="v:starring">西恩·哈里斯</a> / <a href="/celebrity/1025214/" rel="v:starring">安吉拉·贝塞特</a> / <a href="/celebrity/1317680/" rel="v:starring">凡妮莎·柯比</a> / <a href="/celebrity/1000015/" rel="v:starring">米歇尔·莫纳汉</a> / <a href="/celebrity/1049518/" rel="v:starring">韦斯·本特利</a> / <a href="/celebrity/1388344/" rel="v:starring">费雷德里克·施密特</a> / <a href="/celebrity/1000011/" rel="v:starring">亚历克·鲍德温</a> / <a href="/celebrity/1398252/" rel="v:starring">杨亮</a> / <a href="/celebrity/1050483/" rel="v:starring">克里斯托弗·琼勒</a> / <a href="/celebrity/1065337/" rel="v:starring">沃尔夫·布利策</a> / <a href="/celebrity/1398243/" rel="v:starring">拉斐尔·琼勒</a> / <a href="/celebrity/1392979/" rel="v:starring">安德鲁·卡扎纳夫·平</a> / <a href="/celebrity/1398244/" rel="v:starring">克里斯多夫·德·舒瓦西</a> / <a href="/celebrity/1390688/" rel="v:starring">拉裴尔·德普雷</a> / <a href="/celebrity/1392983/" rel="v:starring">让·巴普蒂斯特·菲永</a> / <a href="/celebrity/1398242/" rel="v:starring">马克斯·盖勒</a> / <a href="/celebrity/1392980/" rel="v:starring">奥利维尔·体班德</a> / <a href="/celebrity/1392982/" rel="v:starring">亚历山大·普尔</a> / <a href="/celebrity/1392974/" rel="v:starring">阿利克斯·贝纳泽什</a> / <a href="/celebrity/1036954/" rel="v:starring">乔伊·安沙</a> / <a href="/celebrity/1320241/" rel="v:starring">维利贝·托皮奇</a> / <a href="/celebrity/1392998/" rel="v:starring">格雷厄姆·福克斯</a> / <a href="/celebrity/1367551/" rel="v:starring">卡斯珀·菲利普森</a> / <a href="/celebrity/1392988/" rel="v:starring">菲恩·乔利</a> / <a href="/celebrity/1392999/" rel="v:starring">鲁斯·贝恩</a> / <a href="/celebrity/1392993/" rel="v:starring">奈杰尔·艾伦</a></span></span><br/>
        <span class="pl">类型:</span> <span property="v:genre">动作</span> / <span property="v:genre">惊悚</span> / <span property="v:genre">冒险</span><br/>
        
        <span class="pl">制片国家/地区:</span> 美国<br/>
        <span class="pl">语言:</span> 英语 / 法语<br/>
        <span class="pl">上映日期:</span> <span property="v:initialReleaseDate" content="2018-08-31(中国大陆)">2018-08-31(中国大陆)</span> / <span property="v:initialReleaseDate" content="2018-07-27(美国)">2018-07-27(美国)</span><br/>
        <span class="pl">片长:</span> <span property="v:runtime" content="147">147分钟</span> / 148分钟(中国大陆)<br/>
        <span class="pl">又名:</span> 碟中谍6 / 不可能的任务：全面瓦解(台) / 职业特工队：叛逆之谜(港) / Mission: Impossible 6 / MI6<br/>
        <span class="pl">IMDb链接:</span> <a href="http://www.imdb.com/title/tt4912910" target="_blank" rel="nofollow">tt4912910</a><br>

</div>




                </div>
                    


<div id="interest_sectl">
    <div class="rating_wrap clearbox" rel="v:rating">
        <div class="clearfix">
          <div class="rating_logo ll">豆瓣评分</div>
          <div class="output-btn-wrap rr" style="display:none">
            <img src="https://img3.doubanio.com/f/movie/692e86756648f29457847c5cc5e161d6f6b8aaac/pics/movie/reference.png" />
            <a class="download-output-image" href="#">引用</a>
          </div>
          
          
        </div>
        


<div class="rating_self clearfix" typeof="v:Rating">
    <strong class="ll rating_num" property="v:average">8.3</strong>
    <span property="v:best" content="10.0"></span>
    <div class="rating_right ">
        <div class="ll bigstar bigstar40"></div>
        <div class="rating_sum">
                <a href="collections" class="rating_people"><span property="v:votes">129751</span>人评价</a>
        </div>
    </div>
</div>
<div class="ratings-on-weight">
    
        <div class="item">
        
        <span class="stars5 starstop" title="力荐">
            5星
        </span>
        <div class="power" style="width:41px"></div>
        <span class="rating_per">32.7%</span>
        <br />
        </div>
        <div class="item">
        
        <span class="stars4 starstop" title="推荐">
            4星
        </span>
        <div class="power" style="width:64px"></div>
        <span class="rating_per">50.5%</span>
        <br />
        </div>
        <div class="item">
        
        <span class="stars3 starstop" title="还行">
            3星
        </span>
        <div class="power" style="width:19px"></div>
        <span class="rating_per">15.5%</span>
        <br />
        </div>
        <div class="item">
        
        <span class="stars2 starstop" title="较差">
            2星
        </span>
        <div class="power" style="width:1px"></div>
        <span class="rating_per">1.1%</span>
        <br />
        </div>
        <div class="item">
        
        <span class="stars1 starstop" title="很差">
            1星
        </span>
        <div class="power" style="width:0px"></div>
        <span class="rating_per">0.2%</span>
        <br />
        </div>
</div>

    </div>
        <div class="rating_betterthan">
            好于 <a href="/typerank?type_name=动作&type=5&interval_id=100:90&action=">96% 动作片</a><br/>
            好于 <a href="/typerank?type_name=冒险&type=15&interval_id=100:90&action=">90% 冒险片</a><br/>
        </div>
</div>


                
            </div>
                




<div id="interest_sect_level" class="clearfix">
        
            <a href="https://www.douban.com/reason=collectwish&amp;ck=" rel="nofollow" class="j a_show_login colbutt ll" name="pbtn-26336252-wish">
                <span>想看</span>
            </a>
            <a href="https://www.douban.com/reason=collectcollect&amp;ck=" rel="nofollow" class="j a_show_login colbutt ll" name="pbtn-26336252-collect">
                <span>看过</span>
            </a>
        <div class="ll j a_stars">
            
    
    评价:
    <span id="rating"> <span id="stars" data-solid="https://img3.doubanio.com/f/shire/5a2327c04c0c231bced131ddf3f4467eb80c1c86/pics/rating_icons/star_onmouseover.png" data-hollow="https://img3.doubanio.com/f/shire/2520c01967207a1735171056ec588c8c1257e5f8/pics/rating_icons/star_hollow_hover.png" data-solid-2x="https://img3.doubanio.com/f/shire/7258904022439076d57303c3b06ad195bf1dc41a/pics/rating_icons/star_onmouseover@2x.png" data-hollow-2x="https://img3.doubanio.com/f/shire/95cc2fa733221bb8edd28ad56a7145a5ad33383e/pics/rating_icons/star_hollow_hover@2x.png">

            <a href="https://www.douban.com/register?reason=rate" class="j a_show_login" name="pbtn-26336252-1">
        <img src="https://img3.doubanio.com/f/shire/2520c01967207a1735171056ec588c8c1257e5f8/pics/rating_icons/star_hollow_hover.png" id="star1" width="16" height="16"/></a>
            <a href="https://www.douban.com/register?reason=rate" class="j a_show_login" name="pbtn-26336252-2">
        <img src="https://img3.doubanio.com/f/shire/2520c01967207a1735171056ec588c8c1257e5f8/pics/rating_icons/star_hollow_hover.png" id="star2" width="16" height="16"/></a>
            <a href="https://www.douban.com/register?reason=rate" class="j a_show_login" name="pbtn-26336252-3">
        <img src="https://img3.doubanio.com/f/shire/2520c01967207a1735171056ec588c8c1257e5f8/pics/rating_icons/star_hollow_hover.png" id="star3" width="16" height="16"/></a>
            <a href="https://www.douban.com/register?reason=rate" class="j a_show_login" name="pbtn-26336252-4">
        <img src="https://img3.doubanio.com/f/shire/2520c01967207a1735171056ec588c8c1257e5f8/pics/rating_icons/star_hollow_hover.png" id="star4" width="16" height="16"/></a>
            <a href="https://www.douban.com/register?reason=rate" class="j a_show_login" name="pbtn-26336252-5">
        <img src="https://img3.doubanio.com/f/shire/2520c01967207a1735171056ec588c8c1257e5f8/pics/rating_icons/star_hollow_hover.png" id="star5" width="16" height="16"/></a>
    </span><span id="rateword" class="pl"></span>
    <input id="n_rating" type="hidden" value=""  />
    </span>

        </div>
    

</div>


            


















<div class="gtleft">
    <ul class="ul_subject_menu bicelink color_gray pt6 clearfix">
        
    
        
                <li> 
    <img src="https://img3.doubanio.com/f/shire/cc03d0fcf32b7ce3af7b160a0b85e5e66b47cc42/pics/short-comment.gif" />&nbsp;
        <a onclick="moreurl(this, {from:'mv_sbj_wr_cmnt_login'})" class="j a_show_login" href="https://www.douban.com/register?reason=review" rel="nofollow">写短评</a>
 </li>
                    <li> 
    
    <img src="https://img3.doubanio.com/f/shire/5bbf02b7b5ec12b23e214a580b6f9e481108488c/pics/add-review.gif" />&nbsp;
        <a onclick="moreurl(this, {from:'mv_sbj_wr_rv_login'})" class="j a_show_login" href="https://www.douban.com/register?reason=review" rel="nofollow">写影评</a>
 </li>
                    <li> 
    <img src="https://img3.doubanio.com/f/shire/61cc48ba7c40e0272d46bb93fe0dc514f3b71ec5/pics/add-doulist.gif" />&nbsp;
    <a href="/subject/26336252/questions/ask?from=subject_top">提问题</a>
 </li>
                <li> 
    



 </li>
                <li> 
   

   
    
    <span class="rec" id="电影-26336252">
    <a href= "#"
        data-type="电影"
        data-url="https://movie.douban.com/subject/26336252/"
        data-desc="电影《碟中谍6：全面瓦解 Mission: Impossible - Fallout》 (来自豆瓣) "
        data-title="电影《碟中谍6：全面瓦解 Mission: Impossible - Fallout》 (来自豆瓣) "
        data-pic="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2529365085.jpeg"
        class="bn-sharing ">
        分享到
    </a> &nbsp;&nbsp;
    </span>

    <script>
        if (!window.DoubanShareMenuList) {
            window.DoubanShareMenuList = [];
        }
        var __cache_url = __cache_url || {};

        (function(u){
            if(__cache_url[u]) return;
            __cache_url[u] = true;
            window.DoubanShareIcons = 'https://img3.doubanio.com/f/shire/d15ffd71f3f10a7210448fec5a68eaec66e7f7d0/pics/ic_shares.png';

            var initShareButton = function() {
                $.ajax({url:u,dataType:'script',cache:true});
            };

            if (typeof Do == 'function' && 'ready' in Do) {
                Do(
                    'https://img3.doubanio.com/f/shire/8377b9498330a2e6f056d863987cc7a37eb4d486/css/ui/dialog.css',
                    'https://img3.doubanio.com/f/shire/4ea3216519a6183c7bcd4f7d1a6d4fd57ce1a244/js/ui/dialog.js',
                    'https://img3.doubanio.com/f/movie/c4ab132ff4d3d64a83854c875ea79b8b541faf12/js/movie/lib/qrcode.min.js',
                    initShareButton
                );
            } else if(typeof Douban == 'object' && 'loader' in Douban) {
                Douban.loader.batch(
                    'https://img3.doubanio.com/f/shire/8377b9498330a2e6f056d863987cc7a37eb4d486/css/ui/dialog.css',
                    'https://img3.doubanio.com/f/shire/4ea3216519a6183c7bcd4f7d1a6d4fd57ce1a244/js/ui/dialog.js',
                    'https://img3.doubanio.com/f/movie/c4ab132ff4d3d64a83854c875ea79b8b541faf12/js/movie/lib/qrcode.min.js'
                ).done(initShareButton);
            }

        })('https://img3.doubanio.com/f/movie/32be6727ed3ad8f6c4a417d8a086355c3e7d1d27/js/movie/lib/sharebutton.js');
    </script>


  </li>
            

    </ul>

    <script type="text/javascript">
        $(function(){
            $(".ul_subject_menu li.rec .bn-sharing").bind("click", function(){
                $.get("/blank?sbj_page_click=bn_sharing");
            });
            $(".ul_subject_menu .create_from_menu").bind("click", function(e){
                e.preventDefault();
                var $el = $(this);
                var glRoot = document.getElementById('gallery-topics-selection');
                if (window.has_gallery_topics && glRoot) {
                    // 判断是否有话题
                    glRoot.style.display = 'block';
                } else {
                    location.href = $el.attr('href');
                }
            });
        });
    </script>
</div>




                





<div class="rec-sec">
<span class="rec">
    <script id="movie-share" type="text/x-html-snippet">
        
    <form class="movie-share" action="/j/share" method="POST">
        <div class="clearfix form-bd">
            <div class="input-area">
                <textarea name="text" class="share-text" cols="72" data-mention-api="https://api.douban.com/shuo/in/complete?alt=xd&amp;callback=?"></textarea>
                <input type="hidden" name="target-id" value="26336252">
                <input type="hidden" name="target-type" value="0">
                <input type="hidden" name="title" value="碟中谍6：全面瓦解 Mission: Impossible - Fallout‎ (2018)">
                <input type="hidden" name="desc" value="导演 克里斯托弗·麦奎里 主演 汤姆·克鲁斯 / 亨利·卡维尔 / 美国 / 8.3分(129751评价)">
                <input type="hidden" name="redir" value=""/>
                <div class="mentioned-highlighter"></div>
            </div>

            <div class="info-area">
                    <img class="media" src="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2529365085.webp" />
                <strong>碟中谍6：全面瓦解 Mission: Impossible - Fallout‎ (2018)</strong>
                <p>导演 克里斯托弗·麦奎里 主演 汤姆·克鲁斯 / 亨利·卡维尔 / 美国 / 8.3分(129751评价)</p>
                <p class="error server-error">&nbsp;</p>
            </div>
        </div>
        <div class="form-ft">
            <div class="form-ft-inner">
                



                <span class="avail-num-indicator">140</span>
                <span class="bn-flat">
                    <input type="submit" value="推荐" />
                </span>
            </div>
        </div>
    </form>
    
    <div id="suggest-mention-tmpl" style="display:none;">
        <ul>
            {{#users}}
            <li id="{{uid}}">
              <img src="{{avatar}}">{{{username}}}&nbsp;<span>({{{uid}}})</span>
            </li>
            {{/users}}
        </ul>
    </div>


    </script>

        
        <a href="/accounts/register?reason=recommend"  class="j a_show_login lnk-sharing" share-id="26336252" data-mode="plain" data-name="碟中谍6：全面瓦解 Mission: Impossible - Fallout‎ (2018)" data-type="movie" data-desc="导演 克里斯托弗·麦奎里 主演 汤姆·克鲁斯 / 亨利·卡维尔 / 美国 / 8.3分(129751评价)" data-href="https://movie.douban.com/subject/26336252/" data-image="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2529365085.webp" data-properties="{}" data-redir="" data-text="" data-apikey="" data-curl="" data-count="10" data-object_kind="1002" data-object_id="26336252" data-target_type="rec" data-target_action="0" data-action_props="{&#34;subject_url&#34;:&#34;https:\/\/movie.douban.com\/subject\/26336252\/&#34;,&#34;subject_title&#34;:&#34;碟中谍6：全面瓦解 Mission: Impossible - Fallout‎ (2018)&#34;}">推荐</a>
</span>


</div>






            <script type="text/javascript">
                $(function() {
                    $('.collect_btn', '#interest_sect_level').each(function() {
                        Douban.init_collect_btn(this);
                    });
                    $('html').delegate(".indent .rec-sec .lnk-sharing", "click", function() {
                        moreurl(this, {
                            from : 'mv_sbj_db_share'
                        });
                    });
                });
            </script>
        </div>
            


    <div id="collect_form_26336252"></div>


        



<div class="related-info" style="margin-bottom:-10px;">
    <a name="intro"></a>
    
        
            
            
    <h2>
        <i class="">碟中谍6：全面瓦解的剧情简介</i>
              · · · · · ·
    </h2>

            <div class="indent" id="link-report">
                    
                        <span property="v:summary" class="">
                                　　有时好意会造成恶果，人反而被自己所造成的结果所困扰。伊桑·亨特（汤姆·克鲁斯）和他的IMF团队（亚历克·鲍德温、西蒙·佩吉、文·瑞姆斯）将在最新的电影《碟中谍6：全面瓦解》中再度回归，他们会与观众们熟悉的盟友（丽贝卡·弗格森、米歇尔·莫娜汉）一起与时间赛跑，应对一次任务中出现的意外。亨利·卡维尔、安吉拉·贝塞特和凡妮莎·柯比也将加入本片的演员阵容，电影制片人克里斯托夫·迈考利将会再度担任导演。
                        </span>
                        

            </div>
</div>


    








<div id="celebrities" class="celebrities related-celebrities">

  
    <h2>
        <i class="">碟中谍6：全面瓦解的影人</i>
              · · · · · ·
            <span class="pl">
            (
                <a href="/subject/26336252/celebrities">全部 63</a>
            )
            </span>
    </h2>


  <ul class="celebrities-list from-subject __oneline">
        
    

    
  
  <li class="celebrity">
    

  <a href="https://movie.douban.com/celebrity/1276314/" title="克里斯托弗·麦奎里 " class="">
      <div class="avatar" style="background-image: url(https://img1.doubanio.com/view/celebrity/s_ratio_celebrity/public/p1535912054.09.webp)">
    </div>
  </a>

    <div class="info">
      <span class="name"><a href="https://movie.douban.com/celebrity/1276314/" title="克里斯托弗·麦奎里 " class="name">克里斯托弗·麦奎里 </a></span>

      <span class="role" title="导演">导演</span>

    </div>
  </li>


        
    

    
  
  <li class="celebrity">
    

  <a href="https://movie.douban.com/celebrity/1054435/" title="汤姆·克鲁斯 " class="">
      <div class="avatar" style="background-image: url(https://img1.doubanio.com/view/celebrity/s_ratio_celebrity/public/p567.webp)">
    </div>
  </a>

    <div class="info">
      <span class="name"><a href="https://movie.douban.com/celebrity/1054435/" title="汤姆·克鲁斯 " class="name">汤姆·克鲁斯 </a></span>

      <span class="role" title="饰 伊森·亨特 Ethan Hunt">饰 伊森·亨特 Ethan Hunt</span>

    </div>
  </li>


        
    

    
  
  <li class="celebrity">
    

  <a href="https://movie.douban.com/celebrity/1044713/" title="亨利·卡维尔 " class="">
      <div class="avatar" style="background-image: url(https://img3.doubanio.com/view/celebrity/s_ratio_celebrity/public/p1371934661.95.webp)">
    </div>
  </a>

    <div class="info">
      <span class="name"><a href="https://movie.douban.com/celebrity/1044713/" title="亨利·卡维尔 " class="name">亨利·卡维尔 </a></span>

      <span class="role" title="饰 安格斯·沃克 August Walker">饰 安格斯·沃克 August Walker</span>

    </div>
  </li>


        
    

    
  
  <li class="celebrity">
    

  <a href="https://movie.douban.com/celebrity/1048129/" title="文·瑞姆斯 " class="">
      <div class="avatar" style="background-image: url(https://img3.doubanio.com/view/celebrity/s_ratio_celebrity/public/p8712.webp)">
    </div>
  </a>

    <div class="info">
      <span class="name"><a href="https://movie.douban.com/celebrity/1048129/" title="文·瑞姆斯 " class="name">文·瑞姆斯 </a></span>

      <span class="role" title="饰 路德·史提奇尔 Luther Stickell">饰 路德·史提奇尔 Luther Stickell</span>

    </div>
  </li>


        
    

    
  
  <li class="celebrity">
    

  <a href="https://movie.douban.com/celebrity/1035648/" title="西蒙·佩吉 " class="">
      <div class="avatar" style="background-image: url(https://img3.doubanio.com/view/celebrity/s_ratio_celebrity/public/p8176.webp)">
    </div>
  </a>

    <div class="info">
      <span class="name"><a href="https://movie.douban.com/celebrity/1035648/" title="西蒙·佩吉 " class="name">西蒙·佩吉 </a></span>

      <span class="role" title="饰 班吉·邓恩 Benji Dunn">饰 班吉·邓恩 Benji Dunn</span>

    </div>
  </li>


        
    

    
  
  <li class="celebrity">
    

  <a href="https://movie.douban.com/celebrity/1088314/" title="丽贝卡·弗格森 " class="">
      <div class="avatar" style="background-image: url(https://img3.doubanio.com/view/celebrity/s_ratio_celebrity/public/p1376924506.04.webp)">
    </div>
  </a>

    <div class="info">
      <span class="name"><a href="https://movie.douban.com/celebrity/1088314/" title="丽贝卡·弗格森 " class="name">丽贝卡·弗格森 </a></span>

      <span class="role" title="饰 伊尔莎·浮士德 Ilsa Faust">饰 伊尔莎·浮士德 Ilsa Faust</span>

    </div>
  </li>


  </ul>
</div>


    


<link rel="stylesheet" href="https://img3.doubanio.com/f/verify/16c7e943aee3b1dc6d65f600fcc0f6d62db7dfb4/entry_creator/dist/author_subject/style.css">
<div id="author_subject" class="author-wrapper">
    <div class="loading"></div>
</div>
<script type="text/javascript">
    var answerObj = {
      ISALL: 'False',
      TYPE: 'movie',
      SUBJECT_ID: '26336252',
      USER_ID: 'None'
    }
</script>
<script type="text/javascript" src="https://img3.doubanio.com/f/movie/61252f2f9b35f08b37f69d17dfe48310dd295347/js/movie/lib/react/15.4/bundle.js"></script>
<script type="text/javascript" src="https://img3.doubanio.com/f/verify/ac140ef86262b845d2be7b859e352d8196f3f6d4/entry_creator/dist/author_subject/index.js"></script>


    









    
    <div id="related-pic" class="related-pic">
        
    
    
    <h2>
        <i class="">碟中谍6：全面瓦解的视频和图片</i>
              · · · · · ·
            <span class="pl">
            (
                <a href="https://movie.douban.com/subject/26336252/trailer#trailer">预告片69</a>&nbsp;|&nbsp;<a href="https://movie.douban.com/subject/26336252/trailer#short_video">视频评论1</a>&nbsp;·&nbsp;<a href="/video/create?subject_id=26336252">添加</a>&nbsp;|&nbsp;<a href="https://movie.douban.com/subject/26336252/all_photos">图片609</a>&nbsp;·&nbsp;<a href="https://movie.douban.com/subject/26336252/mupload">添加</a>
            )
            </span>
    </h2>


        <ul class="related-pic-bd  wide_videos">
                <li class="label-trailer">
                    <a class="related-pic-video" href="https://movie.douban.com/trailer/235216/#content" title="预告片" style="background-image:url(https://img3.doubanio.com/img/trailer/medium/2531671156.jpg?1534749628)">
                    </a>
                </li>
                
                <li class="label-short-video">
                    <a class="related-pic-video" href="https://movie.douban.com/video/101045/" title="视频评论" style="background-image:url(https://img3.doubanio.com/view/photo/photo/public/p2532699446.webp?)">
                    </a>
                </li>
                <li>
                    <a href="https://movie.douban.com/photos/photo/2528318164/"><img src="https://img3.doubanio.com/view/photo/sqxs/public/p2528318164.webp" alt="图片" /></a>
                </li>
                <li>
                    <a href="https://movie.douban.com/photos/photo/2511717127/"><img src="https://img1.doubanio.com/view/photo/sqxs/public/p2511717127.webp" alt="图片" /></a>
                </li>
        </ul>
    </div>



    
    



<style type="text/css">
.award li { display: inline; margin-right: 5px }
.awards { margin-bottom: 20px }
.awards h2 { background: none; color: #000; font-size: 14px; padding-bottom: 5px; margin-bottom: 8px; border-bottom: 1px dashed #dddddd }
.awards .year { color: #666666; margin-left: -5px }
.mod { margin-bottom: 25px }
.mod .hd { margin-bottom: 10px }
.mod .hd h2 {margin:24px 0 3px 0}
</style>



    








    <div id="recommendations" class="">
        
        
    <h2>
        <i class="">喜欢这部电影的人也喜欢</i>
              · · · · · ·
    </h2>

        
    
    <div class="recommendations-bd">
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/10727641/?from=subject-page" >
                    <img src="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2263582212.webp" alt="碟中谍5：神秘国度" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/10727641/?from=subject-page" class="" >碟中谍5：神秘国度</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/1578507/?from=subject-page" >
                    <img src="https://img1.doubanio.com/view/photo/s_ratio_poster/public/p792223507.webp" alt="谍影重重3" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/1578507/?from=subject-page" class="" >谍影重重3</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/23761370/?from=subject-page" >
                    <img src="https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2233706697.webp" alt="速度与激情7" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/23761370/?from=subject-page" class="" >速度与激情7</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/26861685/?from=subject-page" >
                    <img src="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2514119443.webp" alt="红海行动" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/26861685/?from=subject-page" class="" >红海行动</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/4920389/?from=subject-page" >
                    <img src="https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2516578307.webp" alt="头号玩家" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/4920389/?from=subject-page" class="" >头号玩家</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/25815034/?from=subject-page" >
                    <img src="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2380677316.webp" alt="湄公河行动" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/25815034/?from=subject-page" class="" >湄公河行动</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/26636712/?from=subject-page" >
                    <img src="https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2529389608.webp" alt="蚁人2：黄蜂女现身" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/26636712/?from=subject-page" class="" >蚁人2：黄蜂女现身</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/24773958/?from=subject-page" >
                    <img src="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2517753454.webp" alt="复仇者联盟3：无限战争" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/24773958/?from=subject-page" class="" >复仇者联盟3：无限战争</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/24405378/?from=subject-page" >
                    <img src="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2231932406.webp" alt="王牌特工：特工学院" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/24405378/?from=subject-page" class="" >王牌特工：特工学院</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/25820460/?from=subject-page" >
                    <img src="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2332503406.webp" alt="美国队长3" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/25820460/?from=subject-page" class="" >美国队长3</a>
            </dd>
        </dl>
    </div>

    </div>



        


<script type="text/x-handlebar-tmpl" id="comment-tmpl">
    <div class="dummy-fold">
        {{#each comments}}
        <div class="comment-item" data-cid="id">
            <div class="comment">
                <h3>
                    <span class="comment-vote">
                            <span class="votes">{{votes}}</span>
                        <input value="{{id}}" type="hidden"/>
                        <a href="javascript:;" class="j {{#if ../if_logined}}a_vote_comment{{else}}a_show_login{{/if}}">有用</a>
                    </span>
                    <span class="comment-info">
                        <a href="{{user.path}}" class="">{{user.name}}</a>
                        {{#if rating}}
                        <span class="allstar{{rating}}0 rating" title="{{rating_word}}"></span>
                        {{/if}}
                        <span>
                            {{time}}
                        </span>
                        <p> {{content_tmpl content}} </p>
                    </span>
            </div>
        </div>
        {{/each}}
    </div>
</script>












    

    <div id="comments-section">
        <div class="mod-hd">
            
        <a class="comment_btn j a_show_login" href="https://www.douban.com/register?reason=review" rel="nofollow">
            <span>我要写短评</span>
        </a>

            
            
    <h2>
        <i class="">碟中谍6：全面瓦解的短评</i>
              · · · · · ·
            <span class="pl">
            (
                <a href="https://movie.douban.com/subject/26336252/comments?status=P">全部 49259 条</a>
            )
            </span>
    </h2>

        </div>
        <div class="mod-bd">
                
    <div class="tab-hd">
        <a id="hot-comments-tab" href="comments" data-id="hot" class="on">热门</a>&nbsp;/&nbsp;
        <a id="new-comments-tab" href="comments?sort=time" data-id="new">最新</a>&nbsp;/&nbsp;
        <a id="following-comments-tab" href="follows_comments" data-id="following"  class="j a_show_login">好友</a>
    </div>

    <div class="tab-bd">
        <div id="hot-comments" class="tab">
            
    
        
        <div class="comment-item" data-cid="1417672930">
            
    
    <div class="comment">
        <h3>
            <span class="comment-vote">
                <span class="votes">4381</span>
                <input value="1417672930" type="hidden"/>
                <a href="javascript:;" class="j a_show_login" onclick="">有用</a>
            </span>
            <span class="comment-info">
                <a href="https://www.douban.com/people/78957039/" class="">维雪魂</a>
                    <span>看过</span>
                    <span class="allstar50 rating" title="力荐"></span>
                <span class="comment-time " title="2018-07-26 19:51:49">
                    2018-07-26
                </span>
            </span>
        </h3>
        <p class="">
            
                <span class="short">只想简单说，在其他谍战动作系列越来越简单粗暴蠢时，碟中谍系列却越战越勇，更加成熟严肃。</span>
        </p>
    </div>

        </div>
        
        <div class="comment-item" data-cid="1418969349">
            
    
    <div class="comment">
        <h3>
            <span class="comment-vote">
                <span class="votes">1436</span>
                <input value="1418969349" type="hidden"/>
                <a href="javascript:;" class="j a_show_login" onclick="">有用</a>
            </span>
            <span class="comment-info">
                <a href="https://www.douban.com/people/2039887/" class="">旺财博士</a>
                    <span>看过</span>
                    <span class="allstar50 rating" title="力荐"></span>
                <span class="comment-time " title="2018-07-28 05:31:12">
                    2018-07-28
                </span>
            </span>
        </h3>
        <p class="">
            
                <span class="short">故事情节老腔老调玩不出什么花头。但不用CG特效、绿幕抠图、替身假演的阿汤哥真的是太拼了：HALO跳伞是真的，厕所格斗是真的，房顶摔断腿是真的，摩托车撞上人甩出去是真的，直升机吊绳上摔下去是真的，克什米尔悬崖徒手攀岩是真的 #职业精神真心服气#</span>
        </p>
    </div>

        </div>
        
        <div class="comment-item" data-cid="1417815712">
            
    
    <div class="comment">
        <h3>
            <span class="comment-vote">
                <span class="votes">639</span>
                <input value="1417815712" type="hidden"/>
                <a href="javascript:;" class="j a_show_login" onclick="">有用</a>
            </span>
            <span class="comment-info">
                <a href="https://www.douban.com/people/NewStart/" class="">Erik Li</a>
                    <span>看过</span>
                    <span class="allstar50 rating" title="力荐"></span>
                <span class="comment-time " title="2018-07-26 22:29:23">
                    2018-07-26
                </span>
            </span>
        </h3>
        <p class="">
            
                <span class="short">这一集，几位女性真是太棒了，包括巴黎的那位小女警。班治和路德的老梗玩得很遛，笑到牙疼。有这样的team，阿汤哥就算七十岁都nothing impossible。阿汤哥的跑跳爬，还有无死角的驾驶技术，是越来越牛逼了。这第6集，不是全面瓦解，是全面巩固。</span>
        </p>
    </div>

        </div>
        
        <div class="comment-item" data-cid="1419157247">
            
    
    <div class="comment">
        <h3>
            <span class="comment-vote">
                <span class="votes">369</span>
                <input value="1419157247" type="hidden"/>
                <a href="javascript:;" class="j a_show_login" onclick="">有用</a>
            </span>
            <span class="comment-info">
                <a href="https://www.douban.com/people/hpjoe/" class="">利岁</a>
                    <span>看过</span>
                    <span class="allstar40 rating" title="推荐"></span>
                <span class="comment-time " title="2018-07-28 11:53:54">
                    2018-07-28
                </span>
            </span>
        </h3>
        <p class="">
            
                <span class="short">单纯视觉上来说，在棚拍人工布光和后期CGI大行其道的近两年，碟6干净不矫揉造作的实景和直白的自然光在IMAX满屏画幅上真有洗眼睛之感</span>
        </p>
    </div>

        </div>
        
        <div class="comment-item" data-cid="1424757555">
            
    
    <div class="comment">
        <h3>
            <span class="comment-vote">
                <span class="votes">275</span>
                <input value="1424757555" type="hidden"/>
                <a href="javascript:;" class="j a_show_login" onclick="">有用</a>
            </span>
            <span class="comment-info">
                <a href="https://www.douban.com/people/qijiuzhiyue/" class="">桃桃淘电影</a>
                    <span>看过</span>
                    <span class="allstar40 rating" title="推荐"></span>
                <span class="comment-time " title="2018-08-02 22:17:13">
                    2018-08-02
                </span>
            </span>
        </h3>
        <p class="">
            
                <span class="short">7，阿汤哥也是拼了老命了。有意思的是，惊险的场面总伴随着搞笑的段落，这倒是很不错的观感。中间那段搏命狂奔，也成为全片最有趣的部分了。超人的角色也不错，尤其是前面，一边凶狠，一边又有些呆萌。丽贝卡·弗格森相比上一部，发挥空间还是有限，略嫌遗憾。作为一部动作大片，碟6也是努力做到更满足观众了，真是高潮一波接一波，尤其还，那么拼命！</span>
        </p>
    </div>

        </div>



                
                &gt; <a href="comments?sort=new_score&status=P" >更多短评49259条</a>
        </div>
        <div id="new-comments" class="tab">
            <div id="normal">
            </div>
            <div class="fold-hd hide">
                <a class="qa" href="/help/opinion#t2-q0" target="_blank">为什么被折叠？</a>
                <a class="btn-unfold" href="#">有一些短评被折叠了</a>
                <div class="qa-tip">
                    评论被折叠，是因为发布这条评论的帐号行为异常。评论仍可以被展开阅读，对发布人的账号不造成其他影响。如果认为有问题，可以<a href="https://help.douban.com/help/ask?category=movie">联系</a>豆瓣电影。
                </div>
            </div>
            <div class="fold-bd">
            </div>
            <span id="total-num"></span>
        </div>
        <div id="following-comments" class="tab">
            
    


        <div class="comment-item">
            你关注的人还没写过短评
        </div>

        </div>
    </div>


            
            
        </div>
    </div>



        

<link rel="stylesheet" href="https://img3.doubanio.com/misc/mixed_static/d04171327f87d16.css">

<section class="topics mod">
    <header>
        <h2>
            碟中谍6：全面瓦解的话题 · · · · · ·
            <span class="pl">( <span class="gallery_topics">全部 <span id="topic-count"></span> 条</span> )</span>
        </h2>
    </header>

    




<section class="subject-topics">
    <div class="topic-guide" id="topic-guide">
        <img class="ic_question" src="//img3.doubanio.com/f/ithildin/b1a3edea3d04805f899e9d77c0bfc0d158df10d5/pics/export/icon_question.png">
        <div class="tip_content">
            <div class="tip_title">什么是话题</div>
            <div class="tip_desc">
                <div>无论是一部作品、一个人，还是一件事，都往往可以衍生出许多不同的话题。将这些话题细分出来，分别进行讨论，会有更多收获。</div>
            </div>
        </div>
        <img class="ic_guide" src="//img3.doubanio.com/f/ithildin/529f46d86bc08f55cd0b1843d0492242ebbd22de/pics/export/icon_guide_arrow.png">
        <img class="ic_close" id="topic-guide-close" src="//img3.doubanio.com/f/ithildin/2eb4ad488cb0854644b23f20b6fa312404429589/pics/export/close@3x.png">
    </div>

    <div id="topic-items"></div>

    <script>
        window.subject_id = 26336252;
        window.join_label_text = '写影评参与';

        window.topic_display_count = 4;
        window.topic_item_display_count = 1;
        window.no_content_fun_call_name = "no_topic";

        window.guideNode = document.getElementById('topic-guide');
        window.guideNodeClose = document.getElementById('topic-guide-close');
    </script>
    
        <link rel="stylesheet" href="https://img3.doubanio.com/f/ithildin/f731c9ea474da58c516290b3a6b1dd1237c07c5e/css/export/subject_topics.css">
        <script src="https://img3.doubanio.com/f/ithildin/d3590fc6ac47b33c804037a1aa7eec49075428c8/js/export/moment-with-locales-only-zh.js"></script>
        <script src="https://img3.doubanio.com/f/ithildin/c600fdbe69e3ffa5a3919c81ae8c8b4140e99a3e/js/export/subject_topics.js"></script>

</section>

    <script>
        function no_topic(){
            $('#content .topics').remove();
        }
    </script>
</section>

<section class="reviews mod movie-content">
    <header>
        <a href="new_review" rel="nofollow" class="create-review comment_btn"
            data-isverify="False"
            data-verify-url="https://www.douban.com/accounts/phone/verify?redir=http://movie.douban.com/subject/26336252/new_review">
            <span>我要写影评</span>
        </a>
        <h2>
            碟中谍6：全面瓦解的影评 · · · · · ·
            <span class="pl">( <a href="reviews">全部 579 条</a> )</span>
        </h2>
    </header>

    

<style>
#gallery-topics-selection {
  position: fixed;
  width: 595px;
  padding: 40px 40px 33px 40px;
  background: #fff;
  border-radius: 10px;
  box-shadow: 0 2px 16px 0 rgba(0, 0, 0, 0.2);
  top: 50%;
  left: 50%;
  -webkit-transform: translate(-50%, -50%);
  transform: translate(-50%, -50%);
  z-index: 9999;
}
#gallery-topics-selection h1 {
  font-size: 18px;
  color: #007722;
  margin-bottom: 36px;
  padding: 0;
  line-height: 28px;
  font-weight: normal;
}
#gallery-topics-selection .gl_topics {
  border-bottom: 1px solid #dfdfdf;
  max-height: 298px;
  overflow-y: scroll;
}
#gallery-topics-selection .topic {
  margin-bottom: 24px;
}
#gallery-topics-selection .topic_name {
  font-size: 15px;
  color: #333;
  margin: 0;
  line-height: inherit;
}
#gallery-topics-selection .topic_meta {
  font-size: 13px;
  color: #999;
}
#gallery-topics-selection .topics_skip {
  display: block;
  cursor: pointer;
  font-size: 16px;
  color: #3377AA;
  text-align: center;
  margin-top: 33px;
}
#gallery-topics-selection .topics_skip:hover {
  background: transparent;
}
#gallery-topics-selection .close_selection {
  position: absolute;
  width: 30px;
  height: 20px;
  top: 46px;
  right: 40px;
  background: #fff;
  color: #999;
  text-align: right;
}
#gallery-topics-selection .close_selection:hover{
  background: #fff;
  color: #999;
}
</style>




        <div class="review_filter">
            <a href="javascript:;;" class="cur" data-sort="">热门</a href="javascript:;;"> /
            <a href="javascript:;;" data-sort="time">最新</a href="javascript:;;"> /
            <a href="javascript:;;" data-sort="follow">好友</a href="javascript:;;">
            
        </div>


        



<div class="review-list  ">
        
    

        
    
    <div xmlns:v="http://rdf.data-vocabulary.org/#" typeof="v:Review" data-cid="9545627">
        <div class="main review-item" id="9545627">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/mbbbbbbbcker/" class="avator">
            <img width="24" height="24" src="https://img3.doubanio.com/icon/u45066005-6.jpg">
        </a>

        <a href="https://www.douban.com/people/mbbbbbbbcker/" property="v:reviewer" class="name">马泽尔法克尔</a>

            <span property="v:rating" class="allstar50 main-title-rating" title="力荐"></span>

        <span property="v:dtreviewed" content="2018-07-28" class="main-meta">2018-07-28 05:15:45</span>


    </header>


            <div class="main-bd">

                <h2><a href="https://movie.douban.com/review/9545627/">电影不需要是3D的，但是导航一定要3D的</a></h2>

                <div id="review_9545627_short" class="review-short" data-rid="9545627">
                    <div class="short-content">
                            <p class="spoiler-tip">这篇影评可能有剧透</p>

                        就算被剧透了，你还是会去电影院看这部电影。不过跟最近的其他影片相比，这一部确实是最怕剧透的，所以请谨慎下翻。 一切纷争由三颗铁蛋引起。准确说来，应该是钚蛋，注意是金字旁，不要因为惯性思维看成土字旁。总之，需要IMF出手的，一定是大规模杀伤性武器。 伊森亨特本来打...

                        &nbsp;(<a href="javascript:;" id="toggle-9545627-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_9545627_full" class="hidden">
                    <div id="review_9545627_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="9545627" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
                        <span id="r-useful_count-9545627">
                                1083
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="9545627" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
                        <span id="r-useless_count-9545627">
                                121
                        </span>
                    </a>
                    <a href="https://movie.douban.com/review/9545627/#comments" class="reply">273回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div xmlns:v="http://rdf.data-vocabulary.org/#" typeof="v:Review" data-cid="9546561">
        <div class="main review-item" id="9546561">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/TheGZMovieGuy/" class="avator">
            <img width="24" height="24" src="https://img3.doubanio.com/icon/u95316262-2.jpg">
        </a>

        <a href="https://www.douban.com/people/TheGZMovieGuy/" property="v:reviewer" class="name">TheGZMovieGuy</a>

            <span property="v:rating" class="allstar50 main-title-rating" title="力荐"></span>

        <span property="v:dtreviewed" content="2018-07-28" class="main-meta">2018-07-28 16:46:43</span>

            <a class="rel-topic" target="_blank" href="//www.douban.com/gallery/topic/《碟中谍6：全面瓦解》与其他同类动作片比有哪些亮点？">#《碟中谍6：全面瓦解》与其他同类动作片比有哪些亮点？</a>

    </header>


            <div class="main-bd">

                <h2><a href="https://movie.douban.com/review/9546561/">严肃影评#44：《碟中谍》系列的极限在哪里？ - 9.5/10</a></h2>

                <div id="review_9546561_short" class="review-short" data-rid="9546561">
                    <div class="short-content">

                        文末有关于高跳低开（HALO Jump）动作场面的详细讨论，有兴趣的朋友可以关注一下！ 转眼22年过去，碟中谍系列已经走到第6部了。22年前俊俏的阿汤哥参与第1部的拍摄时，大概并没有想到这部60年代老旧英剧的翻拍作品可以成为一个系列，更不会想到这个系列作品能有今时今日的成功...

                        &nbsp;(<a href="javascript:;" id="toggle-9546561-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_9546561_full" class="hidden">
                    <div id="review_9546561_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="9546561" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
                        <span id="r-useful_count-9546561">
                                678
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="9546561" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
                        <span id="r-useless_count-9546561">
                                62
                        </span>
                    </a>
                    <a href="https://movie.douban.com/review/9546561/#comments" class="reply">94回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div xmlns:v="http://rdf.data-vocabulary.org/#" typeof="v:Review" data-cid="9552526">
        <div class="main review-item" id="9552526">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/strikeman/" class="avator">
            <img width="24" height="24" src="https://img3.doubanio.com/icon/u96893311-3.jpg">
        </a>

        <a href="https://www.douban.com/people/strikeman/" property="v:reviewer" class="name">Aloneye</a>

            <span property="v:rating" class="allstar50 main-title-rating" title="力荐"></span>

        <span property="v:dtreviewed" content="2018-07-31" class="main-meta">2018-07-31 00:05:18</span>

            <a class="rel-topic" target="_blank" href="//www.douban.com/gallery/topic/《碟中谍6：全面瓦解》与其他同类动作片比有哪些亮点？">#《碟中谍6：全面瓦解》与其他同类动作片比有哪些亮点？</a>

    </header>


            <div class="main-bd">

                <h2><a href="https://movie.douban.com/review/9552526/">《碟中谍6：全面瓦解》，很多你不知道的事！！（宝马M5 系列的两种车型；有腹肌的西蒙......）</a></h2>

                <div id="review_9552526_short" class="review-short" data-rid="9552526">
                    <div class="short-content">

                        你知道阿汤哥的骨折是怎么回事吗？ 你知道“高跳低开”是什么样危险行动吗？ 你知道丽贝卡在拍戏的时候已经怀孕了吗？ 西蒙·佩吉也有八块腹肌了？ 这次片中的宝马轿车，知道是什么型号的吗？ 片中，直升飞机的追逐足够刺激，知道它的型号吗？ ...... 一些数字 这部电影的拍摄...

                        &nbsp;(<a href="javascript:;" id="toggle-9552526-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_9552526_full" class="hidden">
                    <div id="review_9552526_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="9552526" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
                        <span id="r-useful_count-9552526">
                                400
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="9552526" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
                        <span id="r-useless_count-9552526">
                                41
                        </span>
                    </a>
                    <a href="https://movie.douban.com/review/9552526/#comments" class="reply">77回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div xmlns:v="http://rdf.data-vocabulary.org/#" typeof="v:Review" data-cid="9553431">
        <div class="main review-item" id="9553431">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/lingrui1995/" class="avator">
            <img width="24" height="24" src="https://img1.doubanio.com/icon/u63688511-18.jpg">
        </a>

        <a href="https://www.douban.com/people/lingrui1995/" property="v:reviewer" class="name">凌睿</a>

            <span property="v:rating" class="allstar40 main-title-rating" title="推荐"></span>

        <span property="v:dtreviewed" content="2018-07-31" class="main-meta">2018-07-31 12:55:01</span>

            <a class="rel-topic" target="_blank" href="//www.douban.com/gallery/topic/《碟中谍6：全面瓦解》与其他同类动作片比有哪些亮点？">#《碟中谍6：全面瓦解》与其他同类动作片比有哪些亮点？</a>

    </header>


            <div class="main-bd">

                <h2><a href="https://movie.douban.com/review/9553431/">拍摄《碟中谍6》本身就是不可能的任务，但阿汤哥做到了</a></h2>

                <div id="review_9553431_short" class="review-short" data-rid="9553431">
                    <div class="short-content">
                            <p class="spoiler-tip">这篇影评可能有剧透</p>

                        最近，整个网络都在传“《碟中谍6》系列最佳”。单看数据的话也的确如此，截止7月31日MTC 86分，烂番茄新鲜度98%，均分8.4，IMDb 8.5分，均是系列最高分。 《碟中谍6》确实是一部非常优秀的动作片，甚至是今年最佳动作片之一，但我不太喜欢“系列最佳”这种通过贬低前作来捧高...

                        &nbsp;(<a href="javascript:;" id="toggle-9553431-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_9553431_full" class="hidden">
                    <div id="review_9553431_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="9553431" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
                        <span id="r-useful_count-9553431">
                                396
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="9553431" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
                        <span id="r-useless_count-9553431">
                                44
                        </span>
                    </a>
                    <a href="https://movie.douban.com/review/9553431/#comments" class="reply">61回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div xmlns:v="http://rdf.data-vocabulary.org/#" typeof="v:Review" data-cid="9549554">
        <div class="main review-item" id="9549554">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/nggw/" class="avator">
            <img width="24" height="24" src="https://img1.doubanio.com/icon/u2579250-17.jpg">
        </a>

        <a href="https://www.douban.com/people/nggw/" property="v:reviewer" class="name">南瓜国王@似水流年</a>

            <span property="v:rating" class="allstar40 main-title-rating" title="推荐"></span>

        <span property="v:dtreviewed" content="2018-07-29" class="main-meta">2018-07-29 21:04:32</span>


    </header>


            <div class="main-bd">

                <h2><a href="https://movie.douban.com/review/9549554/">阿汤哥：我了个大去的，玩了这么多吧吃鸡，各种救人，你们谁遇到过在天上就需要补血的队友？</a></h2>

                <div id="review_9549554_short" class="review-short" data-rid="9549554">
                    <div class="short-content">

                        警告：以下不但有剧透，还有剧解刨及剧瞎想，未看慎入！ 眼看着007换了一波又一波，伯恩最近也悄无声息，拯救世界的重任最后只好落在年仅六旬的阿汤哥身上 阿汤：哎，谁让我帅啊，没办法… 见过了各种fancy的布置任务的桥段，这次来个朴素的，一个VCR搞定，画面按照COD系列，一...

                        &nbsp;(<a href="javascript:;" id="toggle-9549554-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_9549554_full" class="hidden">
                    <div id="review_9549554_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="9549554" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
                        <span id="r-useful_count-9549554">
                                219
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="9549554" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
                        <span id="r-useless_count-9549554">
                                49
                        </span>
                    </a>
                    <a href="https://movie.douban.com/review/9549554/#comments" class="reply">34回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div xmlns:v="http://rdf.data-vocabulary.org/#" typeof="v:Review" data-cid="9629303">
        <div class="main review-item" id="9629303">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/zishuiyilan/" class="avator">
            <img width="24" height="24" src="https://img3.doubanio.com/icon/u1048262-16.jpg">
        </a>

        <a href="https://www.douban.com/people/zishuiyilan/" property="v:reviewer" class="name">薇羅尼卡</a>

            <span property="v:rating" class="allstar50 main-title-rating" title="力荐"></span>

        <span property="v:dtreviewed" content="2018-09-02" class="main-meta">2018-09-02 00:27:25</span>

            <a class="rel-topic" target="_blank" href="//www.douban.com/gallery/topic/如何看待阿汤哥在《碟中谍6》中的表演？">#如何看待阿汤哥在《碟中谍6》中的表演？</a>

    </header>


            <div class="main-bd">

                <h2><a href="https://movie.douban.com/review/9629303/">阿汤哥这哪里是在演戏，分明是在上纪实真人秀！（暴多拍摄现场图以证清白）</a></h2>

                <div id="review_9629303_short" class="review-short" data-rid="9629303">
                    <div class="short-content">

                        本来没想写东西，结果周五看IMAX场看得我血脉贲张，还想二刷！就是那种你明明知道这个片没啥剧情主要就是让人爽爽爽的，但是爽爽爽和爽爽爽还有分不同的，对不对！这个片可不是什么替身、CGI或者抠图的存在，从第一部到第六部，阿汤哥都从小汤哥变成老汤哥了，依旧生猛！男人的...

                        &nbsp;(<a href="javascript:;" id="toggle-9629303-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_9629303_full" class="hidden">
                    <div id="review_9629303_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="9629303" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
                        <span id="r-useful_count-9629303">
                                117
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="9629303" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
                        <span id="r-useless_count-9629303">
                                10
                        </span>
                    </a>
                    <a href="https://movie.douban.com/review/9629303/#comments" class="reply">18回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div xmlns:v="http://rdf.data-vocabulary.org/#" typeof="v:Review" data-cid="9599640">
        <div class="main review-item" id="9599640">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/157129341/" class="avator">
            <img width="24" height="24" src="https://img1.doubanio.com/icon/u157129341-8.jpg">
        </a>

        <a href="https://www.douban.com/people/157129341/" property="v:reviewer" class="name">CC~</a>

            <span property="v:rating" class="allstar40 main-title-rating" title="推荐"></span>

        <span property="v:dtreviewed" content="2018-08-17" class="main-meta">2018-08-17 23:23:41</span>


    </header>


            <div class="main-bd">

                <h2><a href="https://movie.douban.com/review/9599640/">敬阿汤哥！敬《碟中谍》！</a></h2>

                <div id="review_9599640_short" class="review-short" data-rid="9599640">
                    <div class="short-content">

                        碟中谍从1996年的第一部到现在2018的第六部，转眼就是22年！完全是陪伴了90后的成长，深受其影响，致使如今培养了一大批从小树立间谍梦的青年（包括我在内），虽然长大后我们都会慢慢发现，年少时的间谍梦是那么不切实际，也逐渐将他淡忘于我们的生活，也或许你还会有那么一些...

                        &nbsp;(<a href="javascript:;" id="toggle-9599640-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_9599640_full" class="hidden">
                    <div id="review_9599640_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="9599640" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
                        <span id="r-useful_count-9599640">
                                147
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="9599640" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
                        <span id="r-useless_count-9599640">
                                5
                        </span>
                    </a>
                    <a href="https://movie.douban.com/review/9599640/#comments" class="reply">19回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div xmlns:v="http://rdf.data-vocabulary.org/#" typeof="v:Review" data-cid="9602103">
        <div class="main review-item" id="9602103">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/silence2top/" class="avator">
            <img width="24" height="24" src="https://img3.doubanio.com/icon/u16417472-23.jpg">
        </a>

        <a href="https://www.douban.com/people/silence2top/" property="v:reviewer" class="name">Silence</a>

            <span property="v:rating" class="allstar40 main-title-rating" title="推荐"></span>

        <span property="v:dtreviewed" content="2018-08-19" class="main-meta">2018-08-19 07:43:41</span>


    </header>


            <div class="main-bd">

                <h2><a href="https://movie.douban.com/review/9602103/">建议一刷就看IMAX，二刷好累的</a></h2>

                <div id="review_9602103_short" class="review-short" data-rid="9602103">
                    <div class="short-content">
                            <p class="spoiler-tip">这篇影评可能有剧透</p>

                        本人汤哥20年真爱粉，但不喜动作片所以MI系列都没好好看过。今年初春时候我朋友目击到汤哥在圣保罗大教堂拍外景，急得我那天午休就冒着风雨往圣保罗跑可惜只看到保姆车没看到人。然后电影院里看到的预告片把我一下圈粉，回来恶补了前五部确实帅出天际。 7月底上映后在伦敦一刷P...

                        &nbsp;(<a href="javascript:;" id="toggle-9602103-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_9602103_full" class="hidden">
                    <div id="review_9602103_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="9602103" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
                        <span id="r-useful_count-9602103">
                                45
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="9602103" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
                        <span id="r-useless_count-9602103">
                                5
                        </span>
                    </a>
                    <a href="https://movie.douban.com/review/9602103/#comments" class="reply">19回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div xmlns:v="http://rdf.data-vocabulary.org/#" typeof="v:Review" data-cid="9551875">
        <div class="main review-item" id="9551875">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/1402913/" class="avator">
            <img width="24" height="24" src="https://img3.doubanio.com/icon/u1402913-10.jpg">
        </a>

        <a href="https://www.douban.com/people/1402913/" property="v:reviewer" class="name">大聪看电影</a>

            <span property="v:rating" class="allstar50 main-title-rating" title="力荐"></span>

        <span property="v:dtreviewed" content="2018-07-30" class="main-meta">2018-07-30 20:13:27</span>


    </header>


            <div class="main-bd">

                <h2><a href="https://movie.douban.com/review/9551875/">《碟中谍6：全面瓦解》再次重新定义动作片！</a></h2>

                <div id="review_9551875_short" class="review-short" data-rid="9551875">
                    <div class="short-content">

                        每一部的《碟中谍》，都有阿汤哥固定的招牌动作，这些动作是碟中谍系列在动作片领域江山不倒的基石，而灵魂人物则是阿汤哥不怕死的勇气和行动。 记得第一部《碟中谍1》，是在1996年和父母一起去看的，当时年少的我，就被口香糖炸弹，还有阿汤哥（汤姆克鲁斯）那招牌的密室偷窃...

                        &nbsp;(<a href="javascript:;" id="toggle-9551875-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_9551875_full" class="hidden">
                    <div id="review_9551875_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="9551875" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
                        <span id="r-useful_count-9551875">
                                69
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="9551875" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
                        <span id="r-useless_count-9551875">
                                7
                        </span>
                    </a>
                    <a href="https://movie.douban.com/review/9551875/#comments" class="reply">16回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div xmlns:v="http://rdf.data-vocabulary.org/#" typeof="v:Review" data-cid="9626353">
        <div class="main review-item" id="9626353">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/blueriver007/" class="avator">
            <img width="24" height="24" src="https://img1.doubanio.com/icon/u1325623-628.jpg">
        </a>

        <a href="https://www.douban.com/people/blueriver007/" property="v:reviewer" class="name">小川叔</a>

            <span property="v:rating" class="allstar50 main-title-rating" title="力荐"></span>

        <span property="v:dtreviewed" content="2018-08-31" class="main-meta">2018-08-31 14:06:14</span>

            <a class="rel-topic" target="_blank" href="//www.douban.com/gallery/topic/《碟中谍6：全面瓦解》与其他同类动作片比有哪些亮点？">#《碟中谍6：全面瓦解》与其他同类动作片比有哪些亮点？</a>

    </header>


            <div class="main-bd">

                <h2><a href="https://movie.douban.com/review/9626353/">我推荐《碟中谍6》的六大理由</a></h2>

                <div id="review_9626353_short" class="review-short" data-rid="9626353">
                    <div class="short-content">
                            <p class="spoiler-tip">这篇影评可能有剧透</p>

                        提前看了媒体观影会，看完了只有俩字：激动。 《碟中谍》这个系列，按照现在的说法已经成为了一个大IP，但每一部都会让人有一种期待，又有一种担心。 期待的是，下一次看阿汤哥如何跳战极限，担心的是，这么玩命拍摄……会不会就这样没有了下一部。 所以我的第一个推荐理由就是...

                        &nbsp;(<a href="javascript:;" id="toggle-9626353-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_9626353_full" class="hidden">
                    <div id="review_9626353_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="9626353" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
                        <span id="r-useful_count-9626353">
                                100
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="9626353" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
                        <span id="r-useless_count-9626353">
                                11
                        </span>
                    </a>
                    <a href="https://movie.douban.com/review/9626353/#comments" class="reply">1回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>




    

    

    <script type="text/javascript" src="https://img3.doubanio.com/misc/mixed_static/7d719de715f05719.js"></script>
    <!-- COLLECTED CSS -->
</div>








            <p class="pl">
                &gt;
                <a href="reviews">
                    更多影评579篇
                </a>
            </p>
</section>

<!-- COLLECTED JS -->

    <br/>

        <div class="section-discussion">
                
                <div class="mod-hd">
                        <a class="comment_btn j a_show_login" href="https://www.douban.com/register?reason=review" rel="nofollow"><span>添加新讨论</span></a>
                    
    <h2>
        讨论区
         &nbsp; &middot;&nbsp; &middot;&nbsp; &middot;&nbsp; &middot;&nbsp; &middot;&nbsp; &middot;
    </h2>

                </div>
                
  <table class="olt"><tr><td><td><td><td></tr>
        
        <tr>
          <td class="pl"><a href="https://movie.douban.com/subject/26336252/discussion/615767761/" title="有几个小问题求教？">有几个小问题求教？</a></td>
          <td class="pl"><span>来自</span><a href="https://www.douban.com/people/baohuali/">冰叶日中花</a></td>
          <td class="pl"><span>4 回应</span></td>
          <td class="pl"><span>2018-09-07</span></td>
        </tr>
        
        <tr>
          <td class="pl"><a href="https://movie.douban.com/subject/26336252/discussion/615768935/" title="部长送死想不通">部长送死想不通</a></td>
          <td class="pl"><span>来自</span><a href="https://www.douban.com/people/132957242/">ccthbank_e5o</a></td>
          <td class="pl"><span>6 回应</span></td>
          <td class="pl"><span>2018-09-07</span></td>
        </tr>
        
        <tr>
          <td class="pl"><a href="https://movie.douban.com/subject/26336252/discussion/615768557/" title="大超厕所三连败都是那些？">大超厕所三连败都是那些？</a></td>
          <td class="pl"><span>来自</span><a href="https://www.douban.com/people/2098270/">一个人狂欢</a></td>
          <td class="pl"><span>2 回应</span></td>
          <td class="pl"><span>2018-09-07</span></td>
        </tr>
        
        <tr>
          <td class="pl"><a href="https://movie.douban.com/subject/26336252/discussion/615761451/" title="最后班吉冒充了莱恩，谁冒充了班吉？">最后班吉冒充了莱恩，谁冒充了班吉？</a></td>
          <td class="pl"><span>来自</span><a href="https://www.douban.com/people/love_saiyuki/">仓_仓_仓小颉</a></td>
          <td class="pl"><span>39 回应</span></td>
          <td class="pl"><span>2018-09-07</span></td>
        </tr>
        
        <tr>
          <td class="pl"><a href="https://movie.douban.com/subject/26336252/discussion/615763935/" title="没人吐槽一下所罗门莱恩的战斗力在这部暴涨？">没人吐槽一下所罗门莱恩的战斗力在这部暴涨？</a></td>
          <td class="pl"><span>来自</span><a href="https://www.douban.com/people/156327646/">jaychou</a></td>
          <td class="pl"><span>8 回应</span></td>
          <td class="pl"><span>2018-09-07</span></td>
        </tr>
  </table>

                <p class="pl" align="right">
                    <a href="/subject/26336252/discussion/" rel="nofollow">
                        &gt; 去这部影片的讨论区（全部456条）
                    </a>
                </p>
        </div>

        
    
        
                





<div id="askmatrix">
    <div class="mod-hd">
        <h2>
            关于《碟中谍6：全面瓦解》的问题
            · · · · · ·
            <span class="pl">
                (<a href='https://movie.douban.com/subject/26336252/questions/?from=subject'>
                    全部18个
                </a>)
            </span>
        </h2>


        
    
    <a class='j a_show_login comment_btn'
        href='https://movie.douban.com/subject/26336252/questions/ask/?from=subject'>我来提问</a>

    </div>

    <div class="mod-bd">
        <ul class="">
            <li class="">
                <span class="tit">
                    <a href="https://movie.douban.com/subject/26336252/questions/814206/?from=subject" class="">
                            有人知道厕所激战那一段，几个法国人后来唱的是什么法语歌吗？
                    </a>
                </span>
                <span class="meta">
                    5人回答
                </span>
            </li>
            <li class="">
                <span class="tit">
                    <a href="https://movie.douban.com/subject/26336252/questions/814223/?from=subject" class="">
                            有没觉得白寡妇挺好看的？
                    </a>
                </span>
                <span class="meta">
                    11人回答
                </span>
            </li>
        </ul>

        <p>&gt;
            <a href='https://movie.douban.com/subject/26336252/questions/?from=subject'>
                全部18个问题
            </a>
        </p>

    </div>
</div>



            


    <script type="text/javascript">
        $(function(){if($.browser.msie && $.browser.version == 6.0){
            var $info = $('#info'),
                maxWidth = parseInt($info.css('max-width'));
            if($info.width() > maxWidth) {
                $info.width(maxWidth);
            }
        }});
    </script>


            </div>
            <div class="aside">
                


    









    <!-- douban ad begin -->
    <div id="dale_movie_subject_top_right"></div>
    <div id="dale_movie_subject_top_middle"></div>
    <!-- douban ad end -->

    



<style type="text/css">
    .m4 {margin-bottom:8px; padding-bottom:8px;}
    .movieOnline {background:#FFF6ED; padding:10px; margin-bottom:20px;}
    .movieOnline h2 {margin:0 0 5px;}
    .movieOnline .sitename {line-height:2em; width:160px;}
    .movieOnline td,.movieOnline td a:link,.movieOnline td a:visited{color:#666;}
    .movieOnline td a:hover {color:#fff;}
    .link-bt:link,
    .link-bt:visited,
    .link-bt:hover,
    .link-bt:active {margin:5px 0 0; padding:2px 8px; background:#a8c598; color:#fff; -moz-border-radius: 3px; -webkit-border-radius: 3px; border-radius: 3px; display:inline-block;}
</style>



    







    
    <div class="tags">
        
        
    <h2>
        <i class="">豆瓣成员常用的标签</i>
              · · · · · ·
    </h2>

        <div class="tags-body">
                <a href="/tag/动作" class="">动作</a>
                <a href="/tag/特工" class="">特工</a>
                <a href="/tag/美国" class="">美国</a>
                <a href="/tag/谍战" class="">谍战</a>
                <a href="/tag/2018" class="">2018</a>
                <a href="/tag/间谍" class="">间谍</a>
                <a href="/tag/冒险" class="">冒险</a>
                <a href="/tag/悬疑" class="">悬疑</a>
        </div>
    </div>


    <div id="dale_movie_subject_inner_middle"></div>
    <div id="dale_movie_subject_download_middle"></div>
        








<div id="subject-doulist">
    
    
    <h2>
        <i class="">以下豆列推荐</i>
              · · · · · ·
            <span class="pl">
            (
                <a href="https://movie.douban.com/subject/26336252/doulists">全部</a>
            )
            </span>
    </h2>


    
    <ul>
            <li>
                <a href="https://www.douban.com/doulist/114465/" target="_blank">【北美电影票房总排行榜】</a>
                <span>(荔枝超人)</span>
            </li>
            <li>
                <a href="https://www.douban.com/doulist/30299/" target="_blank">豆瓣电影【口碑榜】2018-02-24更新</a>
                <span>(影志)</span>
            </li>
            <li>
                <a href="https://www.douban.com/doulist/1276354/" target="_blank">2018—2023值得关注的美国影片</a>
                <span>(closer)</span>
            </li>
            <li>
                <a href="https://www.douban.com/doulist/1518184/" target="_blank">IMDB TOP 250 UPDATE 20180825</a>
                <span>(lelechen)</span>
            </li>
            <li>
                <a href="https://www.douban.com/doulist/41896950/" target="_blank">想看的电影太多怕忘了</a>
                <span>(J.D.)</span>
            </li>
    </ul>

</div>

        








<div id="subject-others-interests">
    
    
    <h2>
        <i class="">谁在看这部电影</i>
              · · · · · ·
    </h2>

    
    <ul class="">
    </ul>

    
    <div class="subject-others-interests-ft">
        
            <a href="https://movie.douban.com/subject/26336252/collections">224508人看过</a>
                &nbsp;/&nbsp;
            <a href="https://movie.douban.com/subject/26336252/wishes">274220人想看</a>
    </div>

</div>



    
    

<!-- douban ad begin -->
<div id="dale_movie_subject_middle_right"></div>
<script type="text/javascript">
    (function (global) {
        if(!document.getElementsByClassName) {
            document.getElementsByClassName = function(className) {
                return this.querySelectorAll("." + className);
            };
            Element.prototype.getElementsByClassName = document.getElementsByClassName;

        }
        var articles = global.document.getElementsByClassName('article'),
            asides = global.document.getElementsByClassName('aside');

        if (articles.length > 0 && asides.length > 0 && articles[0].offsetHeight >= asides[0].offsetHeight) {
            (global.DoubanAdSlots = global.DoubanAdSlots || []).push('dale_movie_subject_middle_right');
        }
    })(this);
</script>
<!-- douban ad end -->



    <br/>

    
<p class="pl">订阅碟中谍6：全面瓦解的评论: <br/><span class="feed">
    <a href="https://movie.douban.com/feed/subject/26336252/reviews"> feed: rss 2.0</a></span></p>


            </div>
            <div class="extra">
                
    
<!-- douban ad begin -->
<div id="dale_movie_subject_bottom_super_banner"></div>
<script type="text/javascript">
    (function (global) {
        var body = global.document.body,
            html = global.document.documentElement;

        var height = Math.max(body.scrollHeight, body.offsetHeight, html.clientHeight, html.scrollHeight, html.offsetHeight);
        if (height >= 2000) {
            (global.DoubanAdSlots = global.DoubanAdSlots || []).push('dale_movie_subject_bottom_super_banner');
        }
    })(this);
</script>
<!-- douban ad end -->


            </div>
        </div>
    </div>

        
    <div id="footer">
            <div class="footer-extra"></div>
        
<span id="icp" class="fleft gray-link">
    &copy; 2005－2018 douban.com, all rights reserved 北京豆网科技有限公司
</span>

<a href="https://www.douban.com/hnypt/variformcyst.py" style="display: none;"></a>

<span class="fright">
    <a href="https://www.douban.com/about">关于豆瓣</a>
    · <a href="https://www.douban.com/jobs">在豆瓣工作</a>
    · <a href="https://www.douban.com/about?topic=contactus">联系我们</a>
    · <a href="https://www.douban.com/about?policy=disclaimer">免责声明</a>
    
    · <a href="https://help.douban.com/?app=movie" target="_blank">帮助中心</a>
    · <a href="https://www.douban.com/doubanapp/">移动应用</a>
    · <a href="https://www.douban.com/partner/">豆瓣广告</a>
</span>

    </div>

    </div>
    <script type="text/javascript" src="https://img3.doubanio.com/misc/mixed_static/2b3ad518855276b8.js"></script>
        
        
    <link rel="stylesheet" type="text/css" href="https://img3.doubanio.com/f/shire/8377b9498330a2e6f056d863987cc7a37eb4d486/css/ui/dialog.css" />
    <link rel="stylesheet" type="text/css" href="https://img3.doubanio.com/f/movie/1d829b8605b9e81435b127cbf3d16563aaa51838/css/movie/mod/reg_login_pop.css" />
    <script type="text/javascript" src="https://img3.doubanio.com/f/shire/77323ae72a612bba8b65f845491513ff3329b1bb/js/do.js" data-cfg-autoload="false"></script>
    <script type="text/javascript" src="https://img3.doubanio.com/f/shire/4ea3216519a6183c7bcd4f7d1a6d4fd57ce1a244/js/ui/dialog.js"></script>
    <script type="text/javascript">
        var HTTPS_DB='https://www.douban.com';
var account_pop={open:function(o,e){e?referrer="?referrer="+encodeURIComponent(e):referrer="?referrer="+window.location.href;var n="",i="",t=382;"reg"===o?(n="用户注册",i="https://accounts.douban.com/popup/login?source=movie#popup_register",t=480):"login"===o&&(n="用户登录",i="https://accounts.douban.com/popup/login?source=movie");var r=document.location.protocol+"//"+document.location.hostname,a=dui.Dialog({width:478,title:n,height:t,cls:"account_pop",isHideTitle:!0,modal:!0,content:"<iframe scrolling='no' frameborder='0' width='478' height='"+t+"' src='"+i+"' name='"+r+"'></iframe>"},!0),c=a.node;if(c.undelegate(),c.delegate(".dui-dialog-close","click",function(){var o=$("body");o.find("#login_msk").hide(),o.find(".account_pop").remove()}),$(window).width()<478){var u="";"reg"===o?u=HTTPS_DB+"/accounts/register"+referrer:"login"===o&&(u=HTTPS_DB+"/accounts/login"+referrer),window.location.href=u}else a.open();$(window).bind("message",function(o){"https://accounts.douban.com"===o.originalEvent.origin&&(c.find("iframe").css("height",o.originalEvent.data),c.height(o.originalEvent.data),a.update())})}};Douban&&Douban.init_show_login&&(Douban.init_show_login=function(o){var e=$(o);e.click(function(){var o=e.data("ref")||"";return account_pop.open("login",o),!1})}),Do(function(){$("body").delegate(".pop_register","click",function(o){o.preventDefault();var e=$(this).data("ref")||"";return account_pop.open("reg",e),!1}),$("body").delegate(".pop_login","click",function(o){o.preventDefault();var e=$(this).data("ref")||"";return account_pop.open("login",e),!1})});
    </script>

    
    
    
    




    
<script type="text/javascript">
    (function (global) {
        var newNode = global.document.createElement('script'),
            existingNode = global.document.getElementsByTagName('script')[0],
            adSource = '//erebor.douban.com/',
            userId = '',
            browserId = 'fHkxkxHbMxU',
            criteria = '7:韦斯·本特利|7:犯罪|7:动作|7:克里斯多夫·德·舒瓦西|7:惊悚|7:拉斐尔·琼勒|7:剧情|7:费雷德里克·施密特|7:西恩·哈里斯|7:克里斯托弗·琼勒|7:特工|7:拉裴尔·德普雷|7:2018|7:文·瑞姆斯|7:亨利·卡维尔|7:凡妮莎·柯比|7:汤姆·克鲁斯|7:间谍|7:西蒙·佩吉|7:让·巴普蒂斯特·菲永|7:米歇尔·莫纳汉|7:丽贝卡·弗格森|7:美国|7:谍战|7:悬疑|7:安吉拉·贝塞特|7:冒险|7:克里斯托弗·麦奎里|7:沃尔夫·布利策|7:亚历克·鲍德温|7:安德鲁·卡扎纳夫·平|7:杨亮|3:/subject/26336252/?from=showing',
            preview = '',
            debug = false,
            adSlots = ['dale_movie_subject_top_icon', 'dale_movie_subject_top_right', 'dale_movie_subject_top_middle', 'dale_movie_subject_inner_middle', 'dale_movie_subject_download_middle'];

        global.DoubanAdRequest = {src: adSource, uid: userId, bid: browserId, crtr: criteria, prv: preview, debug: debug};
        global.DoubanAdSlots = (global.DoubanAdSlots || []).concat(adSlots);

        newNode.setAttribute('type', 'text/javascript');
        newNode.setAttribute('src', 'https://img3.doubanio.com/f/adjs/cdc904d1376a43e44bbf399a0aff51973016cd77/ad.release.js');
        newNode.setAttribute('async', true);
        existingNode.parentNode.insertBefore(newNode, existingNode);
    })(this);
</script>











    
  









<script type="text/javascript">
var _paq = _paq || [];
_paq.push(['trackPageView']);
_paq.push(['enableLinkTracking']);
(function() {
    var p=(('https:' == document.location.protocol) ? 'https' : 'http'), u=p+'://fundin.douban.com/';
    _paq.push(['setTrackerUrl', u+'piwik']);
    _paq.push(['setSiteId', '100001']);
    var d=document, g=d.createElement('script'), s=d.getElementsByTagName('script')[0];
    g.type='text/javascript';
    g.defer=true;
    g.async=true;
    g.src=p+'://img3.doubanio.com/dae/fundin/piwik.js';
    s.parentNode.insertBefore(g,s);
})();
</script>

<script type="text/javascript">
var setMethodWithNs = function(namespace) {
  var ns = namespace ? namespace + '.' : ''
    , fn = function(string) {
        if(!ns) {return string}
        return ns + string
      }
  return fn
}

var gaWithNamespace = function(fn, namespace) {
  var method = setMethodWithNs(namespace)
  fn.call(this, method)
}

var _gaq = _gaq || []
  , accounts = [
      { id: 'UA-7019765-1', namespace: 'douban' }
    , { id: 'UA-7019765-19', namespace: '' }
    ]
  , gaInit = function(account) {
      gaWithNamespace(function(method) {
        gaInitFn.call(this, method, account)
      }, account.namespace)
    }
  , gaInitFn = function(method, account) {
      _gaq.push([method('_setAccount'), account.id]);
      _gaq.push([method('_setSampleRate'), '5']);

      
  _gaq.push([method('_addOrganic'), 'google', 'q'])
  _gaq.push([method('_addOrganic'), 'baidu', 'wd'])
  _gaq.push([method('_addOrganic'), 'soso', 'w'])
  _gaq.push([method('_addOrganic'), 'youdao', 'q'])
  _gaq.push([method('_addOrganic'), 'so.360.cn', 'q'])
  _gaq.push([method('_addOrganic'), 'sogou', 'query'])
  if (account.namespace) {
    _gaq.push([method('_addIgnoredOrganic'), '豆瓣'])
    _gaq.push([method('_addIgnoredOrganic'), 'douban'])
    _gaq.push([method('_addIgnoredOrganic'), '豆瓣网'])
    _gaq.push([method('_addIgnoredOrganic'), 'www.douban.com'])
  }

      if (account.namespace === 'douban') {
        _gaq.push([method('_setDomainName'), '.douban.com'])
      }

        _gaq.push([method('_setCustomVar'), 1, 'responsive_view_mode', 'desktop', 3])

        _gaq.push([method('_setCustomVar'), 2, 'login_status', '0', 2]);

      _gaq.push([method('_trackPageview')])
    }

for(var i = 0, l = accounts.length; i < l; i++) {
  var account = accounts[i]
  gaInit(account)
}


;(function() {
    var ga = document.createElement('script');
    ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
    ga.setAttribute('async', 'true');
    document.documentElement.firstChild.appendChild(ga);
})()
</script>

    <!-- brand15-docker-->

  <script>_SPLITTEST=''</script>
</body>

</html>
`
	c.Ctx.WriteString(models.GetMovieDirector(sMovieHtml) + " | ")
	c.Ctx.WriteString(models.GetMovieName(sMovieHtml) + " | ")
	c.Ctx.WriteString(models.GetMovieMainCharactors(sMovieHtml))
}
