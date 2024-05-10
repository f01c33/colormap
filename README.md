# Colormap

Takes a colormap argument, an image and transforms it by going to greyscale and then to the colormap.

### instalation
```bash
go install github.com/f01c33/colormap@latest
```

### Examples
![no colormap](./nature.webp)


```bash
colormap magma nature.webp
```
![magma](./nature.webp_magma_colormap.jpg)

```bash
colormap inferno nature.webp
```
![inferno](./nature.webp_inferno_colormap.jpg)


```bash
colormap viridis nature.webp
```
![viridis](./nature.webp_viridis_colormap.jpg)


```bash
colormap cividis nature.webp
```
![cividis](./nature.webp_cividis_colormap.jpg)


```bash
colormap twilight nature.webp
```
![twilight](./nature.webp_twilight_colormap.jpg)


```bash
colormap turbo nature.webp
```
![turbo](./nature.webp_turbo_colormap.jpg)