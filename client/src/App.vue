<template>
    <div>
        <Name/>
        <Bio/>
        <Tech/>
    </div>
</template>

<script>
import Name from './components/Name'
import Bio from './components/Bio'
import Tech from './components/Tech'

export default {
    name: 'App',
    components: {
        Name,
        Bio,
        Tech
    }
}
</script>

<style>

</style>
