let currentIndex = 0;
let startX = 0;

function showNextImage(nImages) {
    document.getElementById('image-' + currentIndex).style.display = 'none';
    currentIndex++;
    if (currentIndex >= nImages) {
        currentIndex = 0;
    }
    document.getElementById('image-' + currentIndex).style.display = 'block';
}

function showPrevImage(nImages) {
    document.getElementById('image-' + currentIndex).style.display = 'none';
    currentIndex--;
    if (currentIndex < 0) {
        currentIndex = nImages - 1;
    }
    document.getElementById('image-' + currentIndex).style.display = 'block';
}

function updateStatusAndShowNext(image, status, key) {
    const formData = new FormData();
    formData.append('image', image);
    formData.append('status', status);
    formData.append('key', key);
    fetch('/select', {
        method: 'POST',
        body: formData
    }).then(response => {
        if (response.ok) {
            showNextImage();
        }
    });
}

function handleTouchStart(event) {
    startX = event.touches[0].clientX;
}

function handleTouchMove(event) {
    if (!startX) {
        return;
    }
    let endX = event.touches[0].clientX;
    let diffX = startX - endX;
    if (Math.abs(diffX) > 50) {
        document.getElementById('image-' + currentIndex).style.display = 'none';
        if (diffX > 0) {
            // Swipe left
            const image = document.querySelector(`#image-${currentIndex} img`).src.split('/').pop();
            updateStatusAndShowNext(image, 'rejected');
        } else {
            // Swipe right
            const image = document.querySelector(`#image-${currentIndex} img`).src.split('/').pop();
            updateStatusAndShowNext(image, 'selected');
        }
        startX = 0; // reset startX
    }
}

document.querySelector('.swipe-area').addEventListener('touchstart', handleTouchStart);
document.querySelector('.swipe-area').addEventListener('touchmove', handleTouchMove);