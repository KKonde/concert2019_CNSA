
window.onload = function(){
    function getParameterByName(name) {
name = name.replace(/[\[]/, "\\[").replace(/[\]]/, "\\]");
var regex = new RegExp("[\\?&]" + name + "=([^&#]*)"),
    results = regex.exec(location.search);
return results === null ? "" : decodeURIComponent(results[1].replace(/\+/g, " "));
}


a = document.getElementsByClassName('-two')[0];
 b = document.getElementsByClassName('-three')[0];
  setTimeout(function(){
    a.style.opacity = 0;
  },2000);

setTimeout(function(){
  b.style.opacity = 0;
},4000);
setTimeout(function(){
  document.getElementsByClassName('background')[0].classList.add('changebackground');
},1000);
setTimeout(function(){
    document.getElementsByClassName('background')[0].classList.add('zindex3');
   
  },6000);
 

document.getElementsByClassName('check')[0].classList.add('checkup');
var a = getParameterByName("success");
if(a == "false"){
    document.getElementsByClassName('check')[0].innerHTML="신청 기간이 지났습니다.<br />";
    document.getElementsByClassName('check')[0].innerHTML+= "관계자에게 연락을 통해 문의하세요!";
   
    setTimeout(function(){
        
        document.getElementsByClassName('background')[0].classList.add('background-gradation2');
      },6000);
}else{
    setTimeout(function(){
        document.getElementsByClassName('background')[0].classList.add('background-gradation');
       
      },6000);
}
    };