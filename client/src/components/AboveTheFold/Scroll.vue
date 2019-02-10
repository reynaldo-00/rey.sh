<template>
    <div class="mouse-container">
      <div class="mouse"></div>
      <div>
        <span class="eins"></span>
        <span class="zwei"></span>
        <span class="drei"></span>
      </div>
    </div>
</template>

<script>
export default {
    name: "Scroll",
};
</script>

<style lang="scss">
.mouse-container {
  height: 100px;
  position: absolute;
  right: 0;
  bottom: 0;

  span {
    width: 14px;
    height: 15px;
    position: absolute;
    top: 50%;
    left: 50%;
    -webkit-transform: rotate(45deg);
    transform: rotate(45deg);
    border-right: 4px solid white;
    border-bottom: 4px solid white;
    margin: 0 0 0 -8px;
    border-radius: 3px;
  }

  .mouse {
    border: 4px solid white;
    border-radius: 26px;
    position: absolute;
    top: 50%;
    left: 50%;
    width: 30px;
    height: 55px;
    -webkit-backface-visibility: hidden;
    -webkit-transform: translate(-50%, -50%);
    transform: translate(-50%, -50%);
  }

  .mouse::after {
    background-color: white;
    border-radius: 100%;
    content: "";
    position: absolute;
    top: 10px;
    left: 50%;
    margin-left: -2px;
    width: 5px;
    height: 5px;
    opacity: 1;
    -webkit-transform: translateY(0px) scaleY(1) scaleX(1) translateZ(0px);
    transform: translateY(0px) scaleY(1) scaleX(1) translateZ(0px);
    -webkit-animation: scroll 1.5s -1s cubic-bezier(0.68, -0.55, 0.265, 1.55) infinite;
    animation: scroll 1.5s -1s cubic-bezier(0.68, -0.55, 0.265, 1.55) infinite;
  }
  .eins,
  .zwei,
  .drei {
    -webkit-animation: chevron 2.5s infinite;
    animation: chevron 2.5s infinite;
  }

  .eins {
    margin-top: 36px;
    -moz-animation-delay: 0.1s;
    -webkit-animation-delay: 0.1s;
    -webkit-animation-direction: alternate;
  }

  .zwei {
    margin-top: 58px;
    -moz-animation-delay: 0.5s;
    -webkit-animation-delay: 0.5s;
    -webkit-animation-direction: alternate;
  }

  .drei {
    margin-top: 80px;
    -moz-animation-delay: 0.9s;
    -webkit-animation-delay: 0.9s;
    -webkit-animation-direction: alternate;
  }
}
@-webkit-keyframes scroll {
  0%,
  20% {
    -webkit-transform: translateY(0px) scaleY(1) scaleX(1) translateZ(0px);
    transform: translateY(0px) scaleY(1) scaleX(1) translateZ(0px);
  }
  10% {
    -webkit-transform: translateY(0px) scaleY(1.2) scaleX(1.2) translateZ(0px);
    transform: translateY(0px) scaleY(1.2) scaleX(1.2) translateZ(0px);
    opacity: 1;
  }
  100% {
    -webkit-transform: translateY(20px) scaleY(2.5) scaleX(0.5) translateZ(0px);
    transform: translateY(20px) scaleY(2.5) scaleX(0.5) translateZ(0px);
    opacity: 0.01;
  }
}

@keyframes scroll {
  0%,
  20% {
    -webkit-transform: translateY(0px) scaleY(1) scaleX(1) translateZ(0px);
    transform: translateY(0px) scaleY(1) scaleX(1) translateZ(0px);
  }
  10% {
    -webkit-transform: translateY(0px) scaleY(1.2) scaleX(1.2) translateZ(0px);
    transform: translateY(0px) scaleY(1.2) scaleX(1.2) translateZ(0px);
    opacity: 1;
  }
  100% {
    -webkit-transform: translateY(20px) scaleY(2.5) scaleX(0.5) translateZ(0px);
    transform: translateY(20px) scaleY(2.5) scaleX(0.5) translateZ(0px);
    opacity: 0.01;
  }
}

@-webkit-keyframes chevron {
  0% {
    opacity: 0;
  }
  50% {
    opacity: 1;
  }
  100% {
    opacity: 0;
  }
}

@keyframes chevron {
  0% {
    opacity: 0;
  }
  50% {
    opacity: 1;
  }
  100% {
    opacity: 0;
  }
}
</style>
