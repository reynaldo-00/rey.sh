<template>
    <div class="stack-container">
        <span
            v-for="(item, index) in tech"
            :key="index"
            class="tech-item"
        >
            <StyledSVG :logo="item"/>
            <span>{{ firstUpper(item) }}</span>
        </span>
    </div>
</template>

<script>
import StyledSVG from '../../shared/SVG.js';

export default {
    name: 'Stack',
    components: { StyledSVG },
    data: () => ({
        tech: ['react', 'redux', 'styled', 'node', 'knex', 'OAuth'],
    }),
    methods: {
        firstUpper(item) {
            const splitString = item.split("");
            splitString[0] = splitString[0].toUpperCase();
            return splitString.join("");
        }
    }
}
</script>

<style lang="scss" scoped>

.stack-container {
    width: 80%;
    min-width: 380px;
    max-width: 920px;
    display: flex;
    justify-content: space-around;
    align-items: center;
    border: solid 1px white;
    border-radius: 12px;
    padding: 10px;
    align-content: center;
    flex-flow: row wrap;
    margin-top: 60px;
    @media (max-width: 620px) {
        height: 305px;
        min-width: 285px;
        width: 285px;
    }
    @media (max-width: 520px) {
        min-height: 162px;
    }
    .tech-item {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        margin: 0 20px;
        font-size: 16px;
        font-weight: 700;
    }
}

</style>
