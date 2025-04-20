# Eş Zamanlı Dosya İşleyici - ToDo Listesi

## Proje Hakkında
Go dilinde eş zamanlılık kavramlarını öğrenmek için basit bir "Eş Zamanlı Dosya İşleyici" projesi. Bu proje, bir klasördeki tüm metin dosyalarını eş zamanlı olarak işleyip içlerinde belirli kelimeleri arayacak.

## 1. Temel Yapıları Oluşturma
- [ ] Proje dizin yapısını oluştur
  - `concurrent-file-processor/` ana dizini
  - `processor/` paketi
  - `testdata/` klasörü (test için metin dosyaları)
- [ ] Gerekli struct'ları tanımla
  - `Result` struct'ı (dosya işleme sonuçları için)
  - `FileProcessor` struct'ı (ana işlem mantığı için)
- [ ] Ana akış için main.go dosyasını oluştur

## 2. Dosya İşleme Mantığı
- [ ] `listFiles` fonksiyonu: Klasördeki tüm metin dosyalarını listele
- [ ] `processFile` fonksiyonu: Tek bir dosyayı okuyup kelime sayma işlemi yap
- [ ] Basit hata yakalama mantığı ekle

## 3. Eş Zamanlı İşleme Mekanizması
- [ ] Worker havuzu tasarımı
  - [ ] `worker` fonksiyonu: Dosyaları işlemek için goroutine
  - [ ] Yeni goroutine'ler başlatma mantığı
- [ ] Kanal yapısını oluştur
  - [ ] İş kanalı (jobs channel) - dosya isimlerini göndermek için
  - [ ] Sonuç kanalı (results channel) - işleme sonuçlarını almak için
- [ ] `sync.WaitGroup` kullanarak iş tamamlanma kontrolünü ekle

## 4. Sonuçları Toplama ve Raporlama
- [ ] İşlenen tüm dosyaları ve sonuçları haritada (map) topla
- [ ] `sync.Mutex` ile haritaya güvenli erişim sağla
- [ ] Sonuçları düzenli formatta yazdır
- [ ] Toplam işlem süresini ölç ve raporla

## 5. İptal Mekanizması ve Graceful Shutdown
- [ ] `context.Context` kullanarak iptal edilebilir işlemler ekle
- [ ] OS sinyallerini yakala (CTRL+C gibi)
- [ ] Tüm işlemler bitmeden uygulamanın kapanmasını engelle
- [ ] Kaynakları düzgün şekilde temizle

## 6. Ekstra İyileştirmeler (İsteğe Bağlı)
- [ ] Komut satırı argümanları ekle (`flag` paketi ile)
- [ ] İlerleme çubuğu veya işlem durumu göstergesi ekle
- [ ] Bulma eşiklerini ayarlamak için yapılandırma ekle


concurrent-file-processor/
├── main.go               # Ana uygulama
├── processor/
│   ├── processor.go      # FileProcessor yapısı ve metotları
│   └── worker.go         # Worker mantığı
├── testdata/             # Test dosyaları
│   ├── file1.txt
│   ├── file2.txt
│   └── ...
└── README.md             # Proje açıklaması