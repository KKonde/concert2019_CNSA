window.onload = function() {

    //query 가져오기
 function getParameterByName(name) {
  name = name.replace(/[\[]/, "\\[").replace(/[\]]/, "\\]");
  var regex = new RegExp("[\\?&]" + name + "=([^&#]*)"),
      results = regex.exec(location.search);
  return results === null ? "" : decodeURIComponent(results[1].replace(/\+/g, " "));
};
var a = getParameterByName("error");
if(a=="password"){
  document.getElementById('warning').style.opacity=1;
  document.getElementById('warning').textContent="현재 비밀번호가 일치하지 않습니다!";
  document.getElementById('warning').style.color= "red";
 
}
else if (a=="studentNumber"){
  document.getElementById('warning').style.opacity=1;
  document.getElementById('warning').textContent="존재하지 않는 학번입니다.";
  document.getElementById('warning').style.color= "red";
 
} 

  document.getElementById('open').onclick = function() {
    document.getElementById('card').classList.add('on');
    setTimeout(function(){
        document.getElementById('card').classList.add('on3');
    },1200);

    document.getElementById('open').style.opacity = 0;
    setTimeout(function(){

      document.getElementById('table').classList.add('on2');
      document.getElementById('textbox1').classList.add('on2');
      document.getElementById('textbox2').classList.add('on2');
      document.getElementById('login-button').classList.add('on2');
      document.getElementById('login-button').classList.add('testtest');
    },1500);

  }

  document.getElementById('open').onmouseover = function cute(){


      document.getElementById('card').classList.add('up');

  }


  document.getElementById('open').onmouseout = function cute(){
      document.getElementById('card').classList.remove('up');

    }








};
