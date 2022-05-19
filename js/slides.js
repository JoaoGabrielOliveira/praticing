const slides = [
    slide("Natural Environment", 2027, "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vestibulum viverra nunc tincidunt est feugiat tincidunt.", "backgrounds/background1.jpg"),
    slide("Family House", 2025, "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vestibulum viverra nunc tincidunt est feugiat tincidunt.", "backgrounds/background2.jpg"),
    slide("Modern Architecture", 2026, "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vestibulum viverra nunc tincidunt est feugiat tincidunt.", "backgrounds/background3.jpg")
]

function slide(title, date, text, img){
    return {title:title, date:date, text:text, img:img}
}

function showSlide(slide){
    var mainMessage = document.getElementById('main-message');

    mainMessage.getElementsByTagName('h3')[0].textContent = slide.date;
    mainMessage.getElementsByTagName('h1')[0].textContent = slide.title;
    mainMessage.getElementsByTagName('p')[0].textContent = slide.text;
    document.body.style.backgroundImage = `url('${slide.img}')`;
}