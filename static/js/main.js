$(window).on("load", function(){
	// startRaffle()
    $(".pudding-logo").css("transform", " translateX(0)").css("transition", "all 1000ms ease-in-out")
})
   



function setCharAt(str,index,chr) {
    if(index > str.length-1) return str;
    return str.substr(0,index) + chr + str.substr(index+1);
}

function startRaffle(){
	$("#start").css("display", "none")
	$('.text__error').text("");
	$('.text__link a').text("");
	$.get("/raffle", function (result) {
		var raffleWinner = result.winner.Name
	localStorage.setItem("name",raffleWinner)
	localStorage.setItem("company", result.winner.Company.Name)

	})
	
var drum__roll = document.querySelector(".party-drum__roll")
drum__roll.play();


	var characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ`\|'.split('');

var progress404 = 0;
	$('.text__error').attr("data-text", localStorage.getItem("company"))
var total404 = $('.text__error').data('text').length;

var progressLink = 0;
	$('.text__link a').attr("data-text", localStorage.getItem("name"))
var totalLink = $('.text__link a').data('text').length;

var scrambleInterval = setInterval(function() {
	var string404 = $('.text__error').data('text');
	var stringLink = $('.text__link a').data('text');
	
	for(var i = 0; i < total404; i++) {
		if(i >= progress404) {
			string404 = setCharAt(string404, i, characters[Math.round(Math.random() * (characters.length - 1))]);
		} 
	}
	
	for(var i = 0; i < totalLink; i++) {
		if(i >= progressLink) {
			stringLink = setCharAt(stringLink, i, characters[Math.round(Math.random() * (characters.length - 1))]);
		} 
	}
	
	$('.text__error').text(string404);
	$('.text__link a').text(stringLink);	
		

}, 2000 / 60);

setTimeout(function() {
	var revealInterval = setInterval(function() {
		
		if(progress404 < total404) {
			progress404++;
		}else if(progressLink < totalLink) {
			
			setInterval(function(){
			
				progressLink++;
				
				setInterval(function(){
					drum__roll.pause();
					var winner = document.querySelector(".party-drum__winner")
					winner.play();

		
				}, 3000)
			}, 5000)
		
		}else{
			clearInterval(revealInterval);
			clearInterval(scrambleInterval);
		
				$('.modal-container').addClass('active');
				$('.modal-wrapper').addClass('active');
				$("#start").css("display", "flex")
			var name = $('.text__link a').text()
			var department = $('.text__error').text()

			if (department == "                        Altin Pawnshop"){
				$(".modal-wrapper").append(`
			<img src="../static/logo/Altin.svg" alt="pudding-logo">
			`)
			}
			if (department == "           IntegraNet Network Services") {
					$(".modal-wrapper").append(`
		   <img src="../static/logo/Integranet.svg" alt="pudding-logo">
			`)
			}
			if (department == "   IntegrityNet Solutions and Services") {
					$(".modal-wrapper").append(`
			 <img src="../static/logo/integritynet-logo.png" alt="pudding-logo">
			`)
			}
			if (department == "        LuksLofts Hotel and Residences") {
					$(".modal-wrapper").append(`
		  <img src="../static/logo/Luks Lofts Hotel.svg" alt="pudding-logo">
			`)
			}
			if (department == "Share Cafe Restaurant and Event Center") {
					$(".modal-wrapper").append(`
			<img src="../static/logo/Share Cafe.svg" alt="pudding-logo">
			`)
			}
			if (department == "               IntegrityNet Operations") {
					$(".modal-wrapper").append(`
			 <img src="../static/logo/integritynet-logo.png" alt="pudding-logo">
			`)
				}

			$(".modal-wrapper").append(`
			 <h2 class="modal-header">congratulations!</h2>
			 <h1 class="modal-winner">`+ name +`</h1>
			`)
			// setInterval(function () {
			// 	window.location.reload()
			// },10000)
		}
	}, 200);
}, 9000);
}


$(".modal-close").on("click", function(){
	window.location.reload()
})


function Roll(){
	var count = $("#num").val()

	if (count == ""){
		alert("Please Input Number")
	}


	if (count > 15) {
		alert("Maximum of 15 player")
		window.location.reload()
	}
	
	if (count != "" || count <= 15){
		$.ajax({
			url: "/player",
			data: {
				"count": count,
			},
			success: function (res) {
				$(".player-container").empty();
				$.each(res.players, function (k, v) {
					setTimeout(function () {
						console.log(v.Name)
					$(".player-container").append(`
					<div class="col text-focus-in">
						<p class="player">`+v.Name+`</p>
					</div>
					`)	
					  }, k * 1500);
					  var pickPlayer = document.querySelector(".christmas-sound")
					  pickPlayer.play();

					  setTimeout(function(){
						pickPlayer.pause()
					}, 10000)
				});
				
			},
		});
	}
}

function Register(){
	var emp = $("#emps").val()

	if (emp == "") {
		alert("Please Insert Employee Number")
	} else {
		$.ajax({
			url: "/register_emp",
			data: {
				"register": emp,
			},
			success: function () {
				alert("Register Successful")
				window.location.reload()
			},
		});
	}
}


$(document).ready(function () {
	$("#emp").empty()
	$.get("/employee", function (result) {
		$.each(result.employee, function (_, v) {
			$("#emp").append(`
			<option value="`+v.EmployeeNumber+`"> `+v.Name+`</option>
			`)
		})
	})
});
