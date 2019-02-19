<template>
    <section class="section first">
        <div class="background-img"/>
        <Description/>
        <Scroll/>
    </section>
</template>

<script>
import Description from './Description';
import Scroll from './Scroll';

export default {
    name: 'AboveTheFold',
    components: {
        Description,
        Scroll
    }
}
</script>

<style lang="scss">
    section.first {
        padding: 24px;
        position: relative;
        justify-content: center;
        @media (max-width: 620px) {
            display: flex;
            justify-content: center;
            align-items: center;
        }
        .background-img {
            position: fixed;
            width: 100vw;
            height: 100vh;
            max-height: 1080px;
            min-width: 330px;
            left: 0;
            top: 0;
            background: linear-gradient(rgba(0, 0, 0, 0.7),rgba(0, 0, 0, 0.7)), url('../../assets/splash.jpg');
            background-position: center center;
            background-repeat: no-repeat;
            background-size: cover;
        }
    }
</style>
