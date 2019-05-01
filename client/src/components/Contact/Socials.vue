<template>
    <div class="header">
        <a class="item" href="https://github.com/reynld" target="_blank">
            <StyledSVG github/>
        </a>
        <a class="resume" :href="resume" target="__blank">RESUME</a>
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
    },
    data: () => ({
        resume: String
    }),
    beforeMount() {
        this.resume = process.env.VUE_APP_RESUME_URL
    }
}
</script>

<style scoped lang="scss">
.header {
    width: 100%;
    height: 180px;
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    padding: 0;
    a.item {
        width: 60px;
        height: 60px;
        display: flex;
        align-items: center;
        margin: 0 20px;
        cursor: pointer;
        text-decoration: none;
    }
    a.resume {
        text-decoration: none;
        color: white;
        font-size: 22px;
        font-weight: bold;
        letter-spacing: 4px;
        border: solid white 2px;
        border-radius: 16px;
        padding: 5px 20px;
        cursor: pointer;
        user-select: none;
    }
}
</style>
