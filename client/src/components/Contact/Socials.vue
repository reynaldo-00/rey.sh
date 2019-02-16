<template>
    <div class="header">
        <a class="item" href="https://twitter.com/creynl" target="_blank">
            <StyledSVG twitter/>
        </a>
        <a class="item" href="https://codepen.io/reynld/" target="_blank">
            <StyledSVG codepen/>
        </a>
        <a class="item" href="https://github.com/reynld" target="_blank">
            <StyledSVG github/>
        </a>
        <a class="item" href="https://www.linkedin.com/in/reynld/" target="_blank">
            <StyledSVG linkedin/>
        </a>
    </div>
</template>

<script>
import StyledSVG from '../../shared/SVG.js';

export default {
    name: 'Socials',
    components: {
        StyledSVG
    }
}
</script>

<style scoped>
.header {
    width: 100%;
    height: 180px;
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    padding: 0;
}
.item {
    width: 60px;
    height: 60px;
    display: flex;
    align-items: center;
    margin: 0 20px;
    cursor: pointer;
    text-decoration: none;
}
</style>
