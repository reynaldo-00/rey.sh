<template>
    <span class="tech-container">
        <h2>build things with</h2>
        <span
            v-for="(item, index) in tech"
            :key="index"
            class="tech-item"
        >
            <StyledSVG :logo="item"/>
            <span>{{ firstUpper(item) }}</span>
        </span>
    </span>
</template>

<script>
import StyledSVG from './SVG.js';

export default {
    name: 'BuildWith',
    components: { StyledSVG },
    data: () => ({
        tech: ['react', 'redux', 'styled', 'node', 'knex', 'jwt', 'oauth'],
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

<style scoped>

.tech-container {
    width: 90%;
    max-width: 700px;
    position: absolute;
    bottom: 30px;
    display: flex;
    justify-content: space-around;
    align-items: center;
    border: solid 1px white;
    border-radius: 12px;
    padding: 10px;
    align-content: center;
    flex-flow: row wrap;
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

h2 {
    position: absolute;
    top: -70px;
}

</style>
