const slides = [
    slide("Natural Environment", 2027, "Mussum Ipsum, cacilds vidis litro abertis. Sapien in monti palavris qui num significa nadis i pareci latim.Quem num gosta di mé, boa gentis num é.Tá deprimidis, eu conheço uma cachacis que pode alegrar sua vidis.Mais vale um bebadis conhecidiss, que um alcoolatra anonimis.", "backgrounds/background1.jpg"),
    slide("Family House", 2025, "Quem manda na minha terra sou euzis!Suco de cevadiss deixa as pessoas mais interessantis.Nec orci ornare consequat. Praesent lacinia ultrices consectetur. Sed non ipsum felis.Suco de cevadiss, é um leite divinis, qui tem lupuliz, matis, aguis e fermentis.", "backgrounds/background2.jpg"),
    slide("Modern Architecture", 2026, "Si num tem leite então bota uma pinga aí cumpadi!Viva Forevis aptent taciti sociosqu ad litora torquent.Tá deprimidis, eu conheço uma cachacis que pode alegrar sua vidis.Si u mundo tá muito paradis? Toma um mé que o mundo vai girarzis!", "backgrounds/background3.jpg")
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

    getUpdateActualSlide();
}