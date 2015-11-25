$(document).ready(function() {
	var img = $('.image');
    if (img.length > 0) {
        var offset = img.offset();

        function mouse(evt) {
            var center_x = (offset.left) + (img.width() / 2);
            var center_y = (offset.top) + (img.height() / 2);
            var mouse_x = evt.pageX;
            var mouse_y = 1;

            var radians = Math.atan2(center_x - mouse_x, center_y - mouse_y);
            var degree = ((radians * (180 / Math.PI) * -1));
            degree = degree / 2.5;

            img.css('-moz-transform', 'rotate(' + degree + 'deg)');
            img.css('-webkit-transform', 'rotate(' + degree + 'deg)');
            img.css('-o-transform', 'rotate(' + degree + 'deg)');
            img.css('-ms-transform', 'rotate(' + degree + 'deg)');
        }
        $(document).mousemove(mouse);
    }    
});