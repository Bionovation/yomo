var map = null;

let rectangle
let tmprec
const latlngs = []    // 正在绘制的临时矩形的坐标
var drawing = false;  // 正在绘制矩形
var pressing = false; // 鼠标处于按下状态
var modified = false; // 数据已修改
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

var imageList = null
var curImageInfo = null
var curImage = null
var lastMoifiedImage = null
var selectedWsName = null
var workspaceList = null

function init(){
	map = L.map('map', {
		crs: L.CRS.Simple,
		zoomControl:false,
		attributionControl:false,
		minZoom: -3
	});
		
	//getImageList();
	//initImageList();
	
	setmouseHandler();
	setListHandle();
	setClickHandler();
	openWSListDialog()
}

function setClickHandler(){
	$("#workspaceName").click(function(){
		openWSListDialog()
	})
	$("#ws-close").click(function(){
		colseWSListDialog()
	})
	$("#ws-btn-confirm").click(function(){
		if(selectedWsName != null){
			colseWSListDialog()
			openWorkspace(selectedWsName)
		} else{
			return
		}
	})
	$("#ws-btn-cancel").click(function(){
		colseWSListDialog()
	})
	$(".ws-item").click(function(){
		$(".ws-item").css("background-color", "lightgrey")
		$(this).css("background-color", "#ff6666")
		var wsname = $(this).children(":first").text()
		var folder = $(this).children(":last").text()
		selectedWsName = wsname
	})
}

// 提示保存
function noteToSave(){
	if(!modified) return
	var r = confirm("是否保存修改")
	if(r){
		saveChanges()
	}
}

// 更新
function updataImageList(imageList){
	var firstName = null
	var listCtl = document.getElementById('imgList')
	listCtl.innerHTML = ""
	
	for(var i = 0; i < imageList.length; i++){
		var name = imageList[i].Name
		var li = document.createElement('li')
		var span = document.createElement('span')
		span.innerText = name
		li.appendChild(span)
		
		span = document.createElement('span')
		if(imageList[i].Marked){
			span.className = "wancheng iconfont icon-wancheng"
		}else{
			span.className = "wancheng iconfont"
		}
		li.appendChild(span)
		
		li.className = 'imgItem'
		listCtl.appendChild(li)
	}
	
	selectImage(imageList[0])
}

// 绘制mark
function drawMarkList(marks){
	for(var i = 0; i < marks.length; i++){
		mk = marks[i]
		cn = mk.ClassId
		var rc = yolo2leaflet(mk.Rect)
		var rectangle=L.rectangle(rc,{
			color:classClrs[cn],
			fillOpacity:0,
			weight:2,
			className:''+cn,
			attribution: classNames[cn]
		})
		rectangle.addTo(map)
			.bindTooltip(classNames[cn])
			.on('contextmenu', onRightClickItem)
	}
}


// 打开选择工作空间窗口
function openWSListDialog(){
	$("#header").hide()
	$("#main").hide()
	
	selectedWsName = null
	// 从服务器获取ws列表
	// to do ...
	$(".ws-item").css("background-color", "lightgrey")
	$("#wsListDlg").show()
	
	updateWSList()
	
}
// 关闭选择工作空间窗口
function colseWSListDialog(){
	$("#header").show()
	$("#main").show()
	$("#wsListDlg").hide()
}

// 更新工作空间列表
function updateWSList(){
	$("#ws-content").empty()
	$.get("/workspaces", function(data,status){
		if(status == "success"){
			workspaceList = data.data
			drawWorkspaceList(workspaceList)
		}
	})
}
/*
<div class="ws-item"><div>红细胞标注01</div><div>E:\Workdatas\FocusError</div></div>
*/
function drawWorkspaceList(wss){
	$("#ws-content").empty()
	for(var i = 0; i < wss.length; i++){
		$("#ws-content").append('<div class="ws-item"><div>'+wss[i].Name+'</div><div>'+wss[i].Folder+'</div></div>')
	}
	$(".ws-item").click(function(){
		$(".ws-item").css("background-color", "lightgrey")
		$(this).css("background-color", "#ff6666")
		var wsname = $(this).children(":first").text()
		var folder = $(this).children(":last").text()
		selectedWsName = wsname
	})
}

// 打开工作空间
function openWorkspace(name){
	$("#workspaceName").text(name)
	setClassItem()
	initImageList()
}

// 选中image
function selectImage(img){
	$('.imgItem').css('background-color','#F0F8FF')
	var listCtl = document.getElementById('imgList')
	listCtl.childNodes[img.Index].style.backgroundColor = '#ff6666'
	curImage = img.Name
	
	map.eachLayer(function(layer){
		layer.remove()
	})
	
	var bounds = [[0,0], [imgHeight,imgWidth]];
	var image = L.imageOverlay("/"+selectedWsName+"/img/"+img.Name, bounds).addTo(map);
	map.fitBounds(bounds);
	
	initMarkList(img.Name)
}


function initImageList() {
	$.get("/"+selectedWsName+"/imgs",function(data,status){
		if(status == "success"){
			imageList = data.data
			updataImageList(imageList)
		}
	})
}

function initMarkList(imgName) {
	$.get("/"+selectedWsName+"/info/"+imgName, function(data,status){
		if(status == "success"){
			curImageInfo = data.data
			drawMarkList(curImageInfo.Marks)
		}
	})
}

function setListHandle(){
	$('ul#imgList').on('click','li',function(){
		noteToSave(); // 提示保存
		curImage = $(this).text()
		//alert(imgpath)
		map.eachLayer(function(layer){
			layer.remove()
		})
		var bounds = [[0,0], [imgHeight,imgWidth]];
		var image = L.imageOverlay("/"+ selectedWsName + "/img/" + curImage, bounds).addTo(map);
		map.fitBounds(bounds);
		$('.imgItem').css('background-color','#F0F8FF')
		// 列表选中
		$(this).css('background-color','#ff6666')
		modified = false
		
		initMarkList($(this).text())
	})
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
	var curWs = null
	for(var i = 0; i < workspaceList.length; i++){
		if(selectedWsName == workspaceList[i].Name){
			curWs = workspaceList[i]
		}
	}
	if (curWs == null) return;
	
	classNames = curWs.ClassNames
	var hc = document.getElementById("header-content")
	$('.classItem').remove()
	
	for (var i = 0; i < classNames.length; i++){
		var div = document.createElement('div')
		div.className = "classItem"
		div.style.backgroundColor = classClrs[i];
		div.innerText = classNames[i];
		bindClick(div,i)
		hc.appendChild(div)
	}
}

function bindClick(div,i){
	div.onclick = function(){
		setCurClass(i)
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
	modified = true
	lastMoifiedImage = curImage
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
			modified = true
			lastMoifiedImage = curImage
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

function img2latlng(a){
	var e = {}
	e.lng = a.x
	e.lat = imgHeight - a.y
	return e
}

// yolo 转 leaflet
function yolo2leaflet(yrect){
	var cx = imgWidth * yrect[0]
	var cy = imgHeight * yrect[1]
	var w = imgWidth * yrect[2]
	var h = imgHeight * yrect[3]
	var leftbottom = {}
	leftbottom.x = cx - w / 2
	leftbottom.y = cy - h / 2
	var righttop = {}
	righttop.x = cx + w /2
	righttop.y = cy + h / 2
	
	var lb = img2latlng(leftbottom)
	var rt = img2latlng(righttop)
	var leaflet = [[lb.lat, lb.lng], [rt.lat, rt.lng]]
	return leaflet
}


// 保存提交修改
function saveChanges(){
	var mks = []
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
		var mk = {}
		mk.ClassId = parseInt(layer.options.className)
		mk.Rect = pos2
		mks.push(mk)
	});
	
	//console.info(yoloTxt)
	$.ajax({
		url: "/" + selectedWsName + "/mark/" + curImage,
		type: "PUT",
		data: JSON.stringify(mks),
		success: function(data){
			if(data.code == 200){
				modified = false
				if(mks.length > 0){
					markAsLabeled(lastMoifiedImage)
				}
				alert("保存成功")
			}
		}
	})
	
	
	/*$.post("mark/" + curImage, JSON.stringify(mks), function(data, status){
		if(status == "success"){
			if(data.code == 200){
				alert("保存成功")
			}
		}
	})*/
}
// 标记为ok
function markAsLabeled(imgName){
	var imgs = $(".imgItem")
	for(var i = 0; i < imgs.length; i++){
		if (imgs[i].innerText == imgName){
			imgs[i].children[1].className = "iconfont icon-wancheng"
			return
		}
	}
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

