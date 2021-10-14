<template>
  <div class="cropper_model">
    <el-dialog
      title="图片剪裁"
      width="710px"
      class="cropper_model_dlg"
      :visible.sync="dialogVisible"
      append-to-body
      :close-on-click-modal="false"
      :close-on-press-escape="false"
    >
      <div class="cropper_content">
        <div class="cropper" style="text-align: center;">
          <vueCropper
            ref="cropper"
            :img="options.img"
            :outputSize="options.outputSize"
            :outputType="options.outputType"
            :info="options.info"
            :canScale="options.canScale"
            :autoCrop="options.autoCrop"
            :autoCropWidth="options.autoCropWidth"
            :autoCropHeight="options.autoCropHeight"
            :fixed="options.fixed"
            :fixedBox="options.fixedBox"
            :fixedNumber="options.fixedNumber"
            @realTime="previewImg"
          >
          </vueCropper>
          <div class="cropper_btns">
            <el-button
              @click="rotateLeft"
              icon="el-icon-refresh-left"
              size="mini"
              title="左旋转"
            >
            </el-button>
            <el-button
              @click="rotateRight"
              icon="el-icon-refresh-right"
              size="mini"
              title="右旋转"
            >
            </el-button>
            <el-button @click="changeScale(1)" size="mini" title="放大">
              +
            </el-button>
            <el-button @click="changeScale(-1)" size="mini" title="缩小">
              -
            </el-button>
          </div>
        </div>
        <div class="cropper_right">
          <h3>预览</h3>
          <!-- 预览 -->
          <div
            class="cropper_preview"
            :style="{
              width: preview.w + 'px',
              height: preview.h + 'px',
              overflow: 'hidden',
              margin: '5px',
            }"
          >
            <div :style="preview.div">
              <img :src="preview.url" :style="preview.img" alt="" />
            </div>
          </div>
		  <br>
		  <div
		    class="cropper_preview2"
		    :style="{
		      width: preview.w + 'px',
		      height: preview.h + 'px',
		      overflow: 'hidden',
		      margin: '5px',
		    }"
		  >
		    <div :style="preview.div">
		      <img :src="preview.url" :style="preview.img" alt="" />
		    </div>
		  </div>
          <div style="margin-top: 50px;">
            <el-button type="confirmcss" style="width: 150px; margin-left: 45px" @click="uploadImg" :loading="loading">
              确定
            </el-button>
			</div>
			<div style="margin-top: 20px;">
			<el-button type="cancelcss" style="width: 150px; margin-left: 45px" @click="close()" :loading="loading">
			  取消
			</el-button>
          </div>
        </div>
      </div>
      <!-- <div slot="footer" class="dialog-footer">
        <el-button @click="downLoad('blob')">下载</el-button>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="uploadImg" :loading="loading">
          确认
        </el-button>
      </div> -->
    </el-dialog>
  </div>
</template>

<script>
import { VueCropper } from "vue-cropper";
export default {
  components: { VueCropper },
  data() {
    return {
      dialogVisible: false,
      loading: false,
      options: {
        img: "", // 裁剪图片的地址
        outputSize: 1, // 裁剪生成图片的质量
        outputType: "png", // 裁剪生成图片的格式
        info: true, // 裁剪框的大小信息
        canScale: true, // 图片是否允许滚动缩放
        autoCrop: true, // 是否默认生成截图狂
        autoCropWidth: 100, // 默认生成截图框宽度
        autoCropHeight: 100, // 默认生成截图框高度
        fixed: true, // 是否开启截图框宽高固定比例
        fixedNumber: [1, 1], // 截图框的宽高比例
        full: true, // 是否输出原图比例的截图
        fixedBox: true, // 固定截图框大小 不允许改变
        canMove: true, // 上传图片是否可以移动
        canMoveBox: true, // 截图框能否拖动
        original: true, // 上传图片按照原始比例渲染
        centerBox: false, // 截图框是否被限制在图片里面
        high: false, // 是否按照设备的dpr输出等比例图片
        infoTrue: true, // true为展示真实输出图片宽高false展示看到的截图框宽高
        maximgSize: 100, // 限制图片最大宽度和高度
        enlarge: 1, // 图片根据截图框输出比例倍数
        mode: "contain" // 图片默认渲染方式(contain, cover, 100px, 100% auto)
      },
      preview: {}
    };
  },
  methods: {
    open(data) {
      this.options.img = window.URL.createObjectURL(data);
      this.dialogVisible = true;
    },
    close(){
      this.dialogVisible = false;
    },
    // base64转图片文件
    dataURLtoFile(dataurl, filename) {
      let arr = dataurl.split(",");
      let mime = arr[0].match(/:(.*?);/)[1];
      let bstr = atob(arr[1]);
      let len = bstr.length;
      let u8arr = new Uint8Array(len);
      while (len--) {
        u8arr[len] = bstr.charCodeAt(len);
      }
      return new File([u8arr], filename, { type: mime });
    },
    downLoad(type) {
      event.preventDefault();
      const aLink = document.createElement("a");
      if (type === "blob") {
        this.$refs.cropper.getCropBlob(data => {
          let downImg = window.URL.createObjectURL(data);
          aLink.download = "new.png";
          aLink.href = downImg;
          aLink.click();
        });
      } else {
        this.$refs.cropper.getCropData(data => {
          let file = this.dataURLtoFile(data, "test");
          aLink.href = file;
          aLink.click();
        });
      }
    },
    // 左旋转
    rotateLeft() {
      this.$refs.cropper.rotateLeft();
    },
    // 右旋转
    rotateRight() {
      this.$refs.cropper.rotateRight();
    },
    // 放大缩小
    changeScale(num) {
      num = num || 1;
      this.$refs.cropper.changeScale(num);
    },
    // 实时预览
    previewImg(data) {
      this.preview = data;
    },
    // 上传图片
    uploadImg() {
      let self = this;
      self.loading = true;
      this.$refs.cropper.getCropData(data => {
        let file = this.dataURLtoFile(data, "new.png");
        // 生成文件类型
        self.loading = false;
        this.$emit("getFile",file)
      });
    },
    //自定义上传，裁剪后调用
  }
};
</script>

<style lang="less" scoped>
.cropper_model_dlg {
  .cropper_content {
    margin: 0 auto;
    width: 700px;
    height: 450px;
    display: flex;
    justify-content: space-between;
    align-items: flex-start;
  }
  .cropper {
    width: 400px;
    height: 400px;
    background: yellow;
  }
  .cropper_right {
    width: 300px;
    text-align: center;
  }
  .cropper_preview {
    margin: 0 auto;
    display: inline-block;
    border: 1px solid #ddd;
  }
  .cropper_preview2 {
    margin: 0 auto;
    display: inline-block;
    border: 1px solid #ddd;
  	border-radius: 50%;
  }
  .cropper_btns {
    margin-top: 20px;
  }
}
</style>