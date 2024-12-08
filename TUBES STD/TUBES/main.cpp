#include "tubes.h"

int main()
{
    ListPenyewa LP;
    ListBuku LB;
    ListRelasi LR;
    int pilihan = 0;

    createListPenyewa(LP);
    createListBuku(LB);
    createListRelasi(LR);

    infotypePenyewa dataP;
    infotypeBuku dataB;

    adrPenyewa P;
    adrBuku B;
    adrRelasi R;

    string id_penyewa, id_buku, newId_Buku, newId_Penyewa;

    pilihan = menuUtama();
    while (pilihan != 0){
        switch(pilihan){
        case 1:
            system("cls");
            pilihan = menuPenyewa();
            while (pilihan != 0){
                switch(pilihan){
                case 1:
                    system("cls");
                    cout << endl << "============================================================================================" << endl;
                    cout << "|*|                                 TAMBAH DATA PENYEWA                                  |*|" << endl;
                    cout << "============================================================================================" << endl;
                    cout << endl << "Masukkan ID Penyewa : ";
                    cin >> dataP.id_penyewa;

                    validIdPenyewa(LP, dataP.id_penyewa);

                    cout << endl << "Masukkan Nama Penyewa : ";
                    cin >> dataP.nama;
                    cout << endl << "Masukkan NIK Penyewa : ";
                    cin >> dataP.nik;
                    P = createElmPenyewa(dataP);
                    insertLastPenyewa(LP,P);
                    break;

                case 2:
                    system("cls");
                    cout << endl << "============================================================================================" << endl;
                    cout << "|*|                         HAPUS PENYEWA BERDASARKAN ID PENYEWA                         |*|" << endl;
                    cout << "============================================================================================" << endl;
                    cout << endl << "Masukkan ID Penyewa yang Ingin Dihapus : ";
                    cin >> id_penyewa;
                    deletePenyewa(LP, LR, id_penyewa, P, R);
                    break;

                case 3:
                    system("cls");
                    cout << endl << "============================================================================================" << endl;
                    cout << "|*|                          CARI PENYEWA BERDASARKAN ID PENYEWA                         |*|" << endl;
                    cout << "============================================================================================" << endl;
                    cout << endl << "Masukkan ID Penyewa yang Ingin Dicari : ";
                    cin >> id_penyewa;
                    P = findPenyewa(LP, id_penyewa);
                    system("cls");
                    showOneDataPenyewa(P);
                    cout << endl << "ENTER TO NEXT... ";
                    getch();
                    system("cls");
                    break;

                case 4:
                    system("cls");
                    showAllPenyewa(LP);
                    break;
                }
                system("cls");
                pilihan = menuPenyewa();
            }
            break;

        case 2:
            system("cls");
            pilihan = menuBuku();
            while (pilihan != 0){
                switch(pilihan){
                case 1:
                    system("cls");
                    cout << endl << "============================================================================================" << endl;
                    cout << "|*|                                  TAMBAH DATA BUKU                                    |*|" << endl;
                    cout << "============================================================================================" << endl;
                    cout << endl << "Masukkan ID Buku : ";
                    cin >> dataB.id_buku;

                    validIdBuku(LB, dataB.id_buku);

                    cout << endl << "Masukkan Judul Buku : ";
                    cin >> dataB.judul;
                    cout << endl << "Masukkan Penulis Buku : ";
                    cin >> dataB.penulis;
                    cout << endl << "Masukkan Harga Sewa Per Hari : Rp ";
                    cin >> dataB.harga_sewa;
                    B = createElmBuku(dataB);
                    insertFirstBuku(LB,B);
                    break;

                case 2:
                    system("cls");
                    cout << endl << "============================================================================================" << endl;
                    cout << "|*|                            HAPUS BUKU BERDASARKAN ID BUKU                            |*|" << endl;
                    cout << "============================================================================================" << endl;
                    cout << endl << "Masukkan ID Buku yang Ingin Dihapus : ";
                    cin >> id_buku;
                    deleteBuku(LB, LR, id_buku, B, R);
                    break;

                case 3:
                    system("cls");
                    cout << endl << "============================================================================================" << endl;
                    cout << "|*|                            CARI BUKU BERDASARKAN ID BUKU                             |*|" << endl;
                    cout << "============================================================================================" << endl;
                    cout << endl << "Masukkan ID Buku yang Ingin Dicari : ";
                    cin >> id_buku;
                    system("cls");
                    B = findBuku(LB, id_buku);
                    showOneDataBuku(B);
                    cout << endl << "ENTER TO NEXT... ";
                    getch();
                    system("cls");
                    break;

                case 4:
                    system("cls");
                    showAllBuku(LB);
                    break;
                }
                system("cls");
                pilihan = menuBuku();
            }
            break;

        case 3:
            system("cls");
            pilihan = menuRelasi();
            while (pilihan != 0){
                switch(pilihan){
                case 1:
                    system("cls");
                    cout << endl << "============================================================================================" << endl;
                    cout << "|*|                            TAMBAH RELASI PENYEWA DAN BUKU                            |*|" << endl;
                    cout << "============================================================================================" << endl;
                    cout << endl << "Masukkan ID Penyewa : ";
                    cin >> id_penyewa;
                    cout << endl << "Masukkan Id Buku : ";
                    cin >> id_buku;
                    addRelasi(LR, LP, LB, id_penyewa, id_buku);
                    break;

                case 2:
                    system("cls");
                    cout << endl << "============================================================================================" << endl;
                    cout << "|*|                            HAPUS RELASI PENYEWA DAN BUKU                             |*|" << endl;
                    cout << "============================================================================================" << endl;
                    cout << endl << "Masukkan ID Penyewa : ";
                    cin >> id_penyewa;
                    cout << endl << "Masukkan ID Buku : ";
                    cin >> id_buku;
                    deleteRelasi(LR,id_penyewa,id_buku, R);
                    break;

                case 3:
                    system("cls");
                    cout << endl << "============================================================================================" << endl;
                    cout << "|*|                             CARI RELASI PENYEWA DAN BUKU                             |*|" << endl;
                    cout << "============================================================================================" << endl;
                    cout << endl << "Masukkan ID Penyewa : ";
                    cin >> id_penyewa;
                    cout << endl << "Masukkan ID Buku : ";
                    cin >> id_buku;
                    R = findRelasiPenyewaAndBuku(LR, id_penyewa, id_buku);
                    showOneDataRelasi(R);
                    break;

                case 4:
                    system("cls");
                    cout << endl << "============================================================================================" << endl;
                    cout << "|*|                        LIHAT DATA BUKU DARI PENYEWA TERTENTU                         |*|" << endl;
                    cout << "============================================================================================" << endl;
                    cout << endl << "Masukkan ID Penyewa : ";
                    cin >> id_penyewa;
                    system("cls");
                    showBukuFromPenyewa(LR, LP, id_penyewa);
                    cout << endl << "ENTER TO NEXT... ";
                    getch();
                    system("cls");
                    break;

                case 5:
                    system("cls");
                    cout << endl << "============================================================================================" << endl;
                    cout << "|*|                        LIHAT DATA PENYEWA DARI BUKU TERTENTU                         |*|" << endl;
                    cout << "============================================================================================" << endl;
                    cout << endl << "Masukkan ID Buku : ";
                    cin >> id_buku;
                    system("cls");
                    showPenyewaFromBuku(LR, LB, id_buku);
                    cout << endl << "ENTER TO NEXT... ";
                    getch();
                    system("cls");
                    break;

                case 6:
                    system("cls");
                    cout << endl << "============================================================================================" << endl;
                    cout << "|*|                 LIHAT SEMUA DATA PENYEWA DENGAN BUKU YANG DISEWANYA                  |*|" << endl;
                    cout << "============================================================================================" << endl;
                    showAllPenyewaWithBuku(LR, LP);
                    break;

                case 7:
                    system("cls");
                    cout << endl << "============================================================================================" << endl;
                    cout << "|*|                       LIHAT SEMUA DATA BUKU DENGAN PENYEWANYA                        |*|" << endl;
                    cout << "============================================================================================" << endl;
                    showwAllBukuWithPenyewa(LR, LB);
                    break;

                case 8:
                    system("cls");
                    cout << endl << "============================================================================================" << endl;
                    cout << "|*|                       LIHAT JUMLAH BUKU DARI PENYEWA TERTENTU                        |*|" << endl;
                    cout << "============================================================================================" << endl;
                    cout << endl << "Masukkan ID Penyewa : ";
                    cin >> id_penyewa;
                    countBukuDariPenyewa(LR, LP, id_penyewa);
                    break;

                case 9:
                    system("cls");
                    cout << endl << "============================================================================================" << endl;
                    cout << "|*|                       LIHAT JUMLAH PENYEWA DARI BUKU TERTENTU                        |*|" << endl;
                    cout << "============================================================================================" << endl;
                    cout << endl << "Masukkan ID Buku : ";
                    cin >> id_buku;
                    countPenyewaDariBuku(LR, LB, id_buku);
                    break;

                case 10:
                    system("cls");
                    cout << endl << "============================================================================================" << endl;
                    cout << "|*|                         LIHAT JUMLAH BUKU YANG TIDAK DISEWA                          |*|" << endl;
                    cout << "============================================================================================" << endl;
                    countBukuWithNoPenyewa(LR,LB);
                    break;

                case 11:
                    system("cls");
                    cout << endl << "============================================================================================" << endl;
                    cout << "|*|                    LIHAT JUMLAH PENYEWA YANG TIDAK MENYEWA BUKU                      |*|" << endl;
                    cout << "============================================================================================" << endl;
                    countPenyewaWithNoBuku(LR, LP);
                    break;

                case 12:
                    system("cls");
                    cout << endl << "============================================================================================" << endl;
                    cout << "|*|                         EDIT DATA BUKU DARI PENYEWA TERTENTU                         |*|" << endl;
                    cout << "============================================================================================" << endl;
                    cout << endl << "Masukkan ID Penyewa yang Ingin Diedit : ";
                    cin >> id_penyewa;
                    cout << endl << "Masukan ID Buku yang Ingin Diedit : ";
                    cin >> id_buku;
                    cout << endl << "Masukan ID Buku Baru : ";
                    cin >> newId_Buku;
                    editBukuDariPenyewa(LR, LB, id_penyewa, id_buku, newId_Buku);
                    break;

                case 13:
                    system("cls");
                    cout << endl << "============================================================================================" << endl;
                    cout << "|*|                         EDIT DATA PENYEWA DARI BUKU TERTENTU                         |*|" << endl;
                    cout << "============================================================================================" << endl;
                    cout << endl << "Masukkan ID Penyewa yang Ingin Diedit : ";
                    cin >> id_penyewa;
                    cout << endl << "Masukan ID Buku yang Ingin Diedit : ";
                    cin >> id_buku;
                    cout << endl << "Masukan ID Penyewa Baru : ";
                    cin >> newId_Penyewa;
                    editPenyewaDariBuku(LR, LP, id_buku, id_penyewa, newId_Penyewa);
                    break;
                }
                system("cls");
                pilihan = menuRelasi();
            }
            break;
        }
        system("cls");
        pilihan = menuUtama();
    }
    return 0;
}
