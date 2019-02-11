<template>
    <section class="section first">
        <Spotify/>
        <Header/>
        <Description/>
        <Scroll/>
    </section>
</template>

<script>
import Header from './Header';
import Spotify from './Spotify/index'
import Description from './Description';
import Scroll from './Scroll';

export default {
    name: 'AboveTheFold',
    components: {
        Header,
        Description,
        Scroll
    }
}
</script>

<style lang="scss">
    .first {
        height: 100vh;
        min-height: 0px;
        padding: 24px;
    }
</style>
