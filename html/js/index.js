var map = null;

let rectangle
let tmprec
const latlngs = []    // 正在绘制的临时矩形的坐标
var drawing = false;  // 正在绘制矩形
var pressing = false; // 鼠标处于按下状态
var classNo = 0       // 当前要标注的类别id
var classNames = ['RedCell','Platelet','other']
var classClrs = [
	'rgb(230,50,100)',
	'rgb(0,100,255)', 
	'rgb(0,220,0)', 
	"rgb(0,0,0)",
	'rgb(0,220,220)', 
	"rgb(155,44,242)",
	"rgb(240,96,114)", 
	"rgb(172,255,20)"]
	
var imgWidth = 1224
var imgHeight = 1024



function init(){
	map = L.map('map', {
		crs: L.CRS.Simple,
		zoomControl:false,
		attributionControl:false,
		minZoom: -5
	});
	var bounds = [[0,0], [imgHeight,imgWidth]];
	var image = L.imageOverlay('img/2020073001437.jpg', bounds).addTo(map);
	map.fitBounds(bounds);
	
	setmouseHandler();
	setClassItem();
	setCurClass(0);
}

function setCurClass(classId){
	if(classId >= classNames.length){
		alert('error.');
		return;
	}
	classNo = classId;
	var classItems = document.getElementsByClassName('classItem');

	for (var i = 0; i < classItems.length; i++){
		classItems[i].style.border = 'none';
	}
	classItems[classId].style.border = 'solid 2px yellow';
}

function setClassItem(){
	var classCount = classNames.length;
	var classItems = document.getElementsByClassName('classItem');
	if(classItems.length < classCount) alert('class is too many, not supported.');
	
	for (var i = 0; i < classCount; i++){
		classItems[i].style.backgroundColor = classClrs[i];
		classItems[i].innerText = classNames[i];
	}
	for (var i = classItems.length-1; i >= classCount; i--){
		classItems[i].remove();
	}
}

// 绑定鼠标事件
function setmouseHandler(){
	map.on('mousemove', onmouseMove);
	map.on('mousedown', onmouseDown);
	map.on('mouseup', onmouseUp);
}

// 鼠标右键删除
function onRightClickItem(e){
	e.target.remove()
}

function onmouseDown(e){
	if(e.originalEvent.button == 0){
		pressing = true
		//map.dragging.disable();
		
		if(drawing){
			tmprect.remove()
			drawing = false
			latlngs[1]=[e.latlng.lat,e.latlng.lng]
			var rectangle=L.rectangle(latlngs,{
				color:classClrs[classNo],
				fillOpacity:0,
				weight:2,
				className:''+classNo,
				attribution: classNames[classNo]
			})
			rectangle.addTo(map)
				.bindTooltip(classNames[classNo])
				.on('contextmenu', onRightClickItem)
			return
		}
		
		
		if(typeof tmprec != 'undefined'){
			tmprec.remove();
		}
		drawing = true;
		//左上角坐标
        latlngs[0]=[e.latlng.lat,e.latlng.lng];
	}
}

function onmouseUp(e){
	if(e.originalEvent.button == 0){
		pressing = false
	}
}

function onmouseMove(e){
	var a = latlng2img(e.latlng);
	var posStr = 'x:' + Math.round(a.x) + ' y:' + Math.round(a.y);
	document.getElementById('pos').innerHTML = posStr;
	
	if(pressing){ // 鼠标按下的时候移动鼠标
		if(drawing){
			drawing = false
			if (typeof tmprect!='undefined'){
			    tmprect.remove()
			}
		}
	}
	
	if(drawing){
		if (typeof tmprect!='undefined'){
		    tmprect.remove()
		}
		
		latlngs[1]=[e.latlng.lat,e.latlng.lng]
		tmprect=L.rectangle(latlngs,{dashArray:5}).addTo(map)
	}
}


// leaflet 坐标转图像上坐标
function latlng2img(e){
	var a = {}
	a.x = e.lng;
	a.y = imgHeight - e.lat;
	return a;
}

function exportAll(){
	var yoloTxt = ''
	map.eachLayer(function(layer){
		if(layer._url != undefined) return;
		if(layer._bounds._northEast == undefined) return;
		var northEast = latlng2img(layer._bounds._northEast)
		var southWest = latlng2img(layer._bounds._southWest)
		
		var pos = []
		pos.push(southWest.x,northEast.y,northEast.x-southWest.x,southWest.y-northEast.y)
		pos[0] += pos[2]/2
		pos[1] += pos[3]/2
		
		var pos2 = []
		pos2[0] = pos[0] / imgWidth
		pos2[1] = pos[1] / imgHeight
		pos2[2] = pos[2] / imgWidth
		pos2[3] = pos[3] / imgHeight
		
		//var info = layer.options.className + ' ' + JSON.stringify(pos2)
		//console.info(info)
		
		var line = layer.options.className + ' ' 
			 + pos2[0] + ' ' + pos2[1] + ' ' + pos2[2] + ' ' + pos2[3] + '\n'
			//+ pos[0] + ' ' + pos[1] + ' ' + pos[2] + ' ' + pos[3] + '\n'
		yoloTxt += line
	});
	
	console.info(yoloTxt)
}
