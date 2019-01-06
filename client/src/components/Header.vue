<template>
    <div class="container">
        <a class="item" href="https://codepen.io/reynld/" target="_blank">
            <StyledSVG codepen/>
        </a>
        <a class="item" href="https://github.com/reynld" target="_blank">
            <StyledSVG github/>
        </a>
        <a class="item" href="https://www.linkedin.com/in/reynld/" target="_blank">
            <StyledSVG linkedin/>
        </a>
        <!-- <a class="item" href="https://twitter.com/creynl" target="_blank">
            <StyledSVG twitter/>
        </a> -->
    </div>
</template>

<script>
import StyledSVG from './SVG.js';

export default {
    name: 'Header',
    components: {
        StyledSVG
    }
}
</script>

<style scoped>
.container {
    width: 100%;
    height: 120px;
    position: absolute;
    left: 0px;
    top: 60px;
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 0;

}
.item {
    width: 60px;
    height: 60px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    margin: 0 20px;
    cursor: pointer;
    text-decoration: none;
}

</style>
