# 使用轻量级的 Alpine 镜像作为运行环境
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 从 builder 阶段复制编译好的二进制文件
COPY --from=builder /app/myapp .

# 暴露端口
EXPOSE 8080

# 启动应用
ENTRYPOINT ["./myapp"]