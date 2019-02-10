<template>
    <div class="soundbar-container">
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
    </div>
</template>

<script>
export default {
    name: 'SoundBars',
    data: () => ({

    }),
    mounted(){
        let elem = document.querySelectorAll('.bar');
        let count = elem.length;
        let loop = function(){ 
            setTimeout(function(){
                elem.forEach((el, i) => {
                    let height = (Math.random() * 30) + 1;
                    const opacity = .75-( i / count)/2;
                    el.setAttribute("style", `background: rgba(255, 255, 255, ${opacity}); bottom: ${height}px; height: ${height}px`)
                })
                loop()
            }, 300)
        };

        loop();
    }
}
</script>

<style>
    .soundbar-container {
        width: 200px;
        display: flex;
        justify-content: flex-start;
        position: absolute;
        bottom: 0;
        right: 64px;
        z-index: 2;
        transform: rotate(180deg);
    }
    .bar {
        background: #fff;
        width: 2px;
        height: 1px;
        margin-right: 1px;
        transition: all linear 300ms;
        -webkit-box-reflect:below -1px -webkit-gradient(linear, left top, left bottom, from(transparent), color-stop(0.35, transparent), to(black));
    }
</style>
