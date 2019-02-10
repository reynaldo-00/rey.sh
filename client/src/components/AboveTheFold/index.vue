<template>
    <div class="section">
        <Header />
        <Description />
        <Scroll />
    </div>
</template>

<script>
import Header from './Header';
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

</style>
