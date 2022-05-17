var slides = [
    slide("Natural Environment", 2027, "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vestibulum viverra nunc tincidunt est feugiat tincidunt.", "background.jpg")
]

function slide(title, data, text, img){
    return {title:title, data:data, text:text, img:img}
}

function showSlide(slide){
    var mainMessage = document.getElementById('main-message');

    mainMessage.getElementsByTagName('h3')[0].textContent = slide.date;
    mainMessage.getElementsByTagName('h1')[0].textContent = slide.title;
    mainMessage.getElementsByTagName('p')[0].textContent = slide.text;
    document.body.style.backgroundImage = `url('${slide.img}')`;
}