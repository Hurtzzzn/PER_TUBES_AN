#include "tubes.h"

void createListPenyewa(ListPenyewa &LP){
    /*I.S. -
      F.S. Pointer First dari List LP bernilai NIL
    */
    LP.First = NULL;
}

void createListBuku(ListBuku &LB){
    /*I.S. -
      F.S. Pointer First dari List LB benilai NIL
    */
    LB.First = NULL;
    LB.Last = NULL;
}

void createListRelasi(ListRelasi &LR){
    /*I.S. -
      F.S. Pointer First dari List LR benilai NIL
    */
    LR.First = NULL;
}

adrPenyewa createElmPenyewa(infotypePenyewa dataP){
    /*Function akan membuat elemen baru berisi data Penyewa dan
      mengembalikan pointer yang menyimpan alamat dari elemen tersebut
    */
    adrPenyewa P;
    P = new elmPenyewa;
    P->info = dataP;
    P->next = NULL;
    return P;
}

adrBuku createElmBuku(infotypeBuku dataB){
     /*Function akan membuat elemen baru berisi data Buku dan
       mengembalikan pointer yang menyimpan alamat dari elemen tersebut
    */
    adrBuku B = new elmBuku;
    B -> info = dataB;
    B -> next = NULL;
    B -> prev = NULL;
    return B;
}

adrRelasi createElmRelasi(){
      /*Function akan membuat elemen baru dan
       mengembalikan pointer yang menyimpan alamat dari elemen tersebut
    */
    adrRelasi R = new elmRelasi;
    R->nextBuku = NULL;
    R->nextPenyewa = NULL;
    R->nextRelasi = NULL;
    return R;
}

void insertLastPenyewa(ListPenyewa &LP, adrPenyewa P){
    /*I.S. Terdefinisi List LP yang mungkin kosong dan sebuah elemen baru yang alamatnya
           disimpan oleh pointer P
      F.S. Elemen baru telah masuk menjadi elemen paling akhir di List LP
    */
    adrPenyewa Q;
    if (LP.First == NULL){
        LP.First = P;
    }else {
        Q = LP.First;
        while (Q->next != NULL){
            Q = Q->next;
        }
        Q->next = P;
    }
    system("cls");
    cout << endl << "============================================================================================" << endl;
    cout << "|*|                              PENYEWA BERHASIL DITAMBAH!                              |*|" << endl;
    cout << "============================================================================================" << endl;
    cout << endl << "ENTER TO NEXT... ";
    getch();
    system("cls");
}

void insertFirstBuku(ListBuku &LB, adrBuku B){
    /*I.S. Terdefinisi List LB yang mungkin kosong dan sebuah elemen baru yang alamatnya
           disimpan oleh pointer B
      F.S. Elemen baru telah masuk menjadi elemen paling akhir di List LB
    */
    if (LB.First == NULL){
        LB.First = B;
        LB.Last = B;
    }else{
        B -> next = LB.First;
        LB.First -> prev = B;
        LB.First = B;
    }
    system("cls");
    cout << endl << "============================================================================================" << endl;
    cout << "|*|                               BUKU BERHASIL DITAMBAH!!                               |*|" << endl;
    cout << "============================================================================================" << endl;
    cout << endl << "ENTER TO NEXT... ";
    getch();
    system("cls");
}

void addRelasi(ListRelasi &LR, ListPenyewa LP, ListBuku LB, string id_penyewa, string id_buku){
     /*I.S. Terdefinisi ListPenyewa, ListBuku, dan ListRelasi yang mungkin kosong,
            serta id_penyewa dan id_buku yang ingin ditambahkan relasinya.
      F.S. Elemen baru telah masuk menjadi elemen paling akhir di List LR
    */
    adrPenyewa P;
    adrBuku B;
    adrRelasi R, Q;
    system("cls");
    P = findPenyewa(LP, id_penyewa);
    Q = findRelasiPenyewaAndBuku(LR, id_penyewa, id_buku);
    if (Q != NULL){
        cout << endl << "============================================================================================" << endl;
        cout << "|*|                                  RELASI SUDAH ADA!!                                  |*|" << endl;
        cout << "============================================================================================" << endl;
        cout << endl << "ENTER TO NEXT... ";
        getch();
        system("cls");
    } else {
        if (P == NULL){
            cout << endl << "============================================================================================" << endl;
            cout << "|*|                            DATA PENYEWA TIDAK DITEMUKAN!!                            |*|" << endl;
            cout << "============================================================================================" << endl;
            cout << endl << "ENTER TO NEXT... ";
            getch();
            system("cls");
        } else {
            B = findBuku(LB, id_buku);
            if (B == NULL){
                cout << endl << "============================================================================================" << endl;
                cout << "|*|                   DATA PENYEWA DITEMUKAN, TETAPI DATA BUKU TIDAK!!                   |*|" << endl;
                cout << "============================================================================================" << endl;
            } else {
                cout << endl << "============================================================================================" << endl;
                cout << "|*|                           DATA PENYEWA DAN BUKU DITEMUKAN!                           |*|" << endl;
                cout << "============================================================================================" << endl;
                cout << endl << "ENTER TO NEXT... ";
                getch();
                system("cls");
                showOneDataPenyewa(P);
                cout << endl << "ENTER TO NEXT... ";
                getch();
                system("cls");
                showOneDataBuku(B);
                cout << endl << "ENTER TO NEXT... ";
                getch();
                system("cls");
                R = createElmRelasi();
                R->nextPenyewa = P;
                R->nextBuku = B;
                insertLastRelasi(LR, R);
                cout << endl << "============================================================================================" << endl;
                cout << "|*|                              RELASI BERHASIL DITAMBAH!!                              |*|" << endl;
                cout << "============================================================================================" << endl;
            }
            cout << endl << "ENTER TO NEXT... ";
            getch();
            system("cls");
        }
    }
}

void deletePenyewa(ListPenyewa &LP, ListRelasi &LR, string id_penyewa, adrPenyewa &P, adrRelasi &R){
    /*  I.S. Terdefinisi ListPenyewa yang mungkin kosong, atau hanya memiliki 1 elemen
        F.S. Data penyewa tertentu dihapus dari elemen Penyewa
    */
    adrPenyewa Q, precP;
    adrRelasi S, precR;
    Q = findPenyewa(LP, id_penyewa);
    system("cls");
    if (Q == NULL){
        cout << endl << "============================================================================================" << endl;
        cout << "|*|                            DATA PENYEWA TIDAK DITEMUKAN!!                            |*|" << endl;
        cout << "============================================================================================" << endl;
    } else {
        showOneDataPenyewa(Q);
        cout << endl << "ENTER TO NEXT... ";
        getch();
        system("cls");
        S = LR.First;
        while (S != NULL){
            R = S;
            S = S->nextRelasi;
            if (R->nextPenyewa == Q){
                showOneDataRelasi(R);
                if (R == LR.First){
                    deleteFirstRelasi(LR, R);
                } else if (R->nextRelasi == NULL){
                    deleteLastRelasi(LR, R);
                } else {
                    precR = LR.First;
                    while (precR->nextRelasi != R){
                        precR = precR->nextRelasi;
                    }
                    deleteAfterRelasi(LR, precR, R);
                }
                delete(R);
                cout << endl << "============================================================================================" << endl;
                cout << "|*|                            DATA RELASI BERHASIL DIHAPUS!!                            |*|" << endl;
                cout << "============================================================================================" << endl;
                cout << endl << "ENTER TO NEXT... ";
                getch();
                system("cls");
            }
        }
        if (Q == LP.First){
            deleteFirstPenyewa(LP, P);
        } else if (Q->next == NULL){
            deleteLastPenyewa(LP, P);
        } else {
            precP = LP.First;
            while (precP->next != Q){
                precP = precP->next;
            }
            deleteAfterPenyewa(LP, precP, P);
        }
        delete(P);
        cout << endl << "============================================================================================" << endl;
        cout << "|*|                            DATA PENYEWA BERHASIL DIHAPUS!                            |*|" << endl;
        cout << "============================================================================================" << endl;
    }
    cout << endl << "ENTER TO NEXT... ";
    getch();
    system("cls");
}

void deleteBuku(ListBuku &LB, ListRelasi &LR, string id_buku, adrBuku &B, adrRelasi &R){
    /*  I.S. Terdefinisi ListBuku yang mungkin kosong, atau hanya memiliki 1 elemen
        F.S. Data buku tertentu dihapus dari elemen Buku
    */

    adrBuku Q, precB;
    adrRelasi S, precR;

    system("cls");
    Q = findBuku(LB, id_buku);
    if (Q == NULL){
        cout << endl << "============================================================================================" << endl;
        cout << "|*|                              DATA BUKU TIDAK DITEMUKAN!                              |*|" << endl;
        cout << "============================================================================================" << endl;
    } else{
        showOneDataBuku(Q);
        cout << endl << "ENTER TO NEXT... ";
        getch();
        system("cls");
        S = LR.First;
        while (S != NULL)
        {
            R = S;
            S = S -> nextRelasi;
            if (R -> nextBuku == Q){
                showOneDataRelasi(R);
                if (R == LR.First)
                {
                    deleteFirstRelasi(LR, R);
                }else if (R -> nextRelasi == NULL)
                {
                    deleteLastRelasi(LR, R);
                }else {
                    precR = LR.First;
                    while (precR -> nextRelasi != R)
                    {
                        precR = precR -> nextRelasi;
                    }
                    deleteAfterRelasi(LR, precR, R);
                }
                delete(R);
                cout << endl << "============================================================================================" << endl;
                cout << "|*|                            DATA RELASI BERHASIL DIHAPUS!!                            |*|" << endl;
                cout << "============================================================================================" << endl;
                cout << endl << "ENTER TO NEXT... ";
                getch();
                system("cls");
            }
       }
        if (Q == LB.First){
            deleteFirstBuku(LB, B);
        } else if (Q->next == NULL){
            deleteLastBuku(LB, B);
        } else {
            precB = LB.First;
            while (precB->next != Q){
                precB = precB->next;
            }
            deleteAfterBuku(LB, precB, B);
        }
        delete(B);
        cout << endl << "============================================================================================" << endl;
        cout << "|*|                             DATA BUKU BERHASIL DIHAPUS!!                             |*|" << endl;
        cout << "============================================================================================" << endl;
    }
    cout << endl << "ENTER TO NEXT... ";
    getch();
    system("cls");
}

void deleteRelasi(ListRelasi &LR, string id_penyewa, string id_buku, adrRelasi &R){
    /*  I.S. Terdefinisi ListRelasi yang mungkin kosong, atau hanya memiliki 1 elemen,
            serta id_penyewa dan id_buku untuk mencari relasi keduanya.
        F.S. Data Relasi tertentu dihapus dari ListRelasi LR dan alamatnya disimpan di pointer R jika relasi ditemukan,
             jika tidak maka akan tercetak di layar "DATA RELASI TIDAK DITEMUKAN".
    */

    adrRelasi Q, prec;

    system("cls");
    if (LR.First == NULL){
        cout << endl << "============================================================================================" << endl;
        cout << "|*|                                 LIST RELASI KOSONG!!                                 |*|" << endl;
        cout << "============================================================================================" << endl;
    } else {
        Q = findRelasiPenyewaAndBuku(LR, id_penyewa, id_buku);
        if (Q == NULL){
            cout << endl << "============================================================================================" << endl;
            cout << "|*|                             DATA RELASI TIDAK DITEMUKAN!                             |*|" << endl;
            cout << "============================================================================================" << endl;
        } else {
            showOneDataRelasi(Q);
            if (Q == LR.First) {
                deleteFirstRelasi(LR, R);
            } else if (Q->nextRelasi == NULL){
                deleteLastRelasi(LR, R);
            } else {
                prec = LR.First;
                while (prec->nextRelasi != Q){
                    prec = prec->nextRelasi;
                }
                deleteAfterRelasi(LR, prec, R);
            }
            delete(R);
            cout << endl << "============================================================================================" << endl;
            cout << "|*|                            DATA RELASI BERHASIL DIHAPUS!!                            |*|" << endl;
            cout << "============================================================================================" << endl;
        }
    }
    cout << endl << "ENTER TO NEXT... ";
    getch();
    system("cls");
}

adrPenyewa findPenyewa(ListPenyewa LP, string id_penyewa){
     /*fungsi ini akan mencari id_penyewa pada ListPenyewa, jika ditemukan maka akan
       dikembalikan alamatnya, jika tidak ditemukan maka akan dikembalikan NIL
    */
    adrPenyewa P;
    bool ketemu = false;
    P = LP.First;
    while (P != NULL && !ketemu){
        ketemu = P->info.id_penyewa == id_penyewa;
            if (!ketemu){
            P = P->next;
        }
    }
    return P;
}

adrBuku findBuku(ListBuku LB, string id_buku){
    /*fungsi ini akan mencari id_buku pada ListBuku, jika ditemukan maka akan
      dikembalikan alamatnya, jika tidak ditemukan maka akan dikembalikan NIL
    */
    adrBuku B = LB.First;
    while (B != NULL){
        if (B -> info.id_buku == id_buku){
            return B;
        }
        B = B -> next;
    }
    return NULL;
}

adrRelasi findRelasiPenyewaAndBuku(ListRelasi LR, string id_penyewa, string id_buku){
    /*fungsi ini akan mencari relasi dari id_penyewa dan id_buku tertentu pada ListRelasi,
      jika ditemukan maka akan dikembalikan alamatnya, jika tidak ditemukan maka akan dikembalikan NIL
    */
    adrRelasi R;
    bool ketemu = false;
    R = LR.First;
    while (R != NULL && !ketemu){
        ketemu = R->nextPenyewa->info.id_penyewa == id_penyewa && R->nextBuku->info.id_buku == id_buku;
        if (!ketemu){
            R = R->nextRelasi;
        }
    }
    return R;
}

void showAllPenyewa(ListPenyewa LP){
    /*I.S. Terdefinisi ListPenyewa yang mungkin kosong
      F.S. Data Penyewa tampil ke layar
    */
    adrPenyewa P;
    system("cls");
    if (LP.First == NULL){
        cout << endl << "============================================================================================" << endl;
        cout << "|*|                                 LIST PENYEWA KOSONG!                                 |*|" << endl;
        cout << "============================================================================================" << endl;
    }else{
        cout << endl << "============================================================================================" << endl;
        cout << "|*|                                     DATA PENYEWA                                     |*|" << endl;
        cout << "============================================================================================" << endl;

        P = LP.First;
        while (P != NULL){
            cout << endl << "ID      : " << P->info.id_penyewa << endl;
            cout << "Nama    : " << P->info.nama << endl;
            cout << "NIK     : " << P->info.nik << endl;
            P = P->next;
        }
        cout << endl << "============================================================================================" << endl;
    }
    cout << endl << "ENTER TO NEXT... ";
    getch();
    system("cls");
}

void showAllBuku(ListBuku LB){
    /*I.S. Terdefinisi ListBuku yang mungkin kosong
      F.S. Data Buku tampil ke layar
    */
    adrBuku B = LB.First;

    system("cls");
    if (LB.First == NULL){
        cout << endl << "============================================================================================" << endl;
        cout << "|*|                                  LIST BUKU KOSONG!!                                  |*|" << endl;
        cout << "============================================================================================" << endl;
    } else {
        cout << endl << "============================================================================================" << endl;
        cout << "|*|                                      DATA BUKU                                       |*|" << endl;
        cout << "============================================================================================" << endl;

        while (B != NULL){
            cout << endl << "ID                  : " << B->info.id_buku << endl;
            cout << "Judul               : " << B->info.judul << endl;
            cout << "Penulis             : " << B->info.penulis << endl;
            cout << "Harga Sewa Per Hari : Rp " << B->info.harga_sewa << endl;
            B = B->next;
        }
        cout << endl << "============================================================================================" << endl;
    }
    cout << endl << "ENTER TO NEXT... ";
    getch();
    system("cls");
}

void showBukuFromPenyewa(ListRelasi LR, ListPenyewa LP, string id_penyewa){
    /* I.S. Terdefinisi List Relasi (LR) yang mungkin kosong dan string id_penyewa.
       F.S. Semua buku yang berelasi dengan penyewa tertentu ditampilkan ke layar.
            Jika tidak ditemukan, ditampilkan pesan sesuai kondisi.
    */

    adrRelasi R;
    adrPenyewa P;
    bool notKetemu = true;
    bool visitedBuku = false;

    if (LR.First == NULL){
        cout << endl << "============================================================================================" << endl;
        cout << "|*|                                 LIST RELASI KOSONG!!                                 |*|" << endl;
        cout << "============================================================================================" << endl;
    } else {
        P = findPenyewa(LP, id_penyewa);

        showOneDataPenyewa(P);

        if (P != NULL){
            R = LR.First;
            while (R != NULL){
                if (R->nextPenyewa == P){
                    if (!visitedBuku){
                        cout << endl << "============================================================================================" << endl;
                        cout << "|*|                                  DATA BUKU YANG DISEWA                               |*|" << endl;
                        cout << "============================================================================================" << endl;
                        visitedBuku = true;
                    }
                    cout << endl << "ID                  : " << R->nextBuku->info.id_buku << endl;
                    cout << "Judul               : " << R->nextBuku->info.judul << endl;
                    cout << "Penulis             : " << R->nextBuku->info.penulis << endl;
                    cout << "Harga Sewa Per Hari : Rp " << R->nextBuku->info.harga_sewa << endl;
                    notKetemu = false;
                }
                R = R->nextRelasi;
            }

            if (notKetemu){
                cout << endl << "============================================================================================" << endl;
                cout << "|*|                               TIDAK ADA BUKU YANG DISEWA!                            |*|" << endl;
                cout << "============================================================================================" << endl;
            } else {
                cout << endl << "============================================================================================" << endl;
            }
        }
    }
}

void showPenyewaFromBuku(ListRelasi LR, ListBuku LB, string id_buku){
    /* I.S. Terdefinisi List Relasi (LR) yang mungkin kosong dan string id_buku.
       F.S. Semua penyewa yang berelasi dengan buku tertentu ditampilkan ke layar.
            Jika tidak ditemukan, ditampilkan pesan sesuai kondisi.
    */


    adrRelasi R = LR.First;
    adrBuku B;
    bool found = false;
    bool visitedPenyewa = false;

    if (LR.First == NULL){
        cout << endl << "============================================================================================" << endl;
        cout << "|*|                                 LIST RELASI KOSONG!!                                 |*|" << endl;
        cout << "============================================================================================" << endl;
    } else {
        B = findBuku(LB, id_buku);

        showOneDataBuku(B);

        if (B != NULL){
            while (R != NULL){
                if (R -> nextBuku == B){
                    if (!visitedPenyewa){
                        cout << endl << "============================================================================================" << endl;
                        cout << "|*|                            DATA YANG MENYEWA BUKU TERSEBUT                           |*|" << endl;
                        cout << "============================================================================================" << endl;
                        visitedPenyewa = true;
                    }
                    cout << endl << "ID      : " << R->nextPenyewa->info.id_penyewa << endl;
                    cout << "Nama    : " << R->nextPenyewa->info.nama << endl;
                    cout << "NIK     : " << R->nextPenyewa->info.nik << endl;
                    found = true;
                }
                R = R -> nextRelasi;
            }

            if (!found){
                cout << endl << "============================================================================================" << endl;
                cout << "|*|                         TIDAK ADA YANG MENYEWA BUKU TERSEBUT!                        |*|" << endl;
                cout << "============================================================================================" << endl;
            } else {
                cout << endl << "============================================================================================" << endl;
            }
        }
    }
}

void showAllPenyewaWithBuku(ListRelasi LR, ListPenyewa LP){
    /* I.S. Terdefinisi List Relasi (LR) dan ListPenyewa (LP)
       F.S. Semua data Penyewa yang berelasi dengan buku ditampilkan ke layar beserta buku yang disewanya.
            Jika tidak ada relasi, data penyewa tidak ditampilkan.
    */

    adrPenyewa P;
    adrRelasi R;
    bool visitedPenyewa = false;

    system("cls");
    if (LR.First == NULL){
        cout << endl << "============================================================================================" << endl;
        cout << "|*|                                 LIST RELASI KOSONG!!                                 |*|" << endl;
        cout << "============================================================================================" << endl;
        cout << endl << "ENTER TO NEXT... ";
        getch();
        system("cls");
    } else {
        P = LP.First;
        while (P != NULL){
            R = LR.First;
            while (R != NULL){
                if (R->nextPenyewa == P){
                    if (!visitedPenyewa){
                        cout << endl << "============================================================================================" << endl;
                        cout << "|*|                                     DATA PENYEWA                                     |*|" << endl;
                        cout << "============================================================================================" << endl;
                        cout << endl << "ID      : " << P->info.id_penyewa << endl;
                        cout << "Nama    : " << P->info.nama << endl;
                        cout << "NIK     : " << P->info.nik << endl;
                        cout << endl << "============================================================================================" << endl;

                        visitedPenyewa = true;
                        cout << endl << "============================================================================================" << endl;
                        cout << "|*|                                  DATA BUKU YANG DISEWA                               |*|" << endl;
                        cout << "============================================================================================" << endl;
                    }
                    cout << endl << "ID                  : " << R->nextBuku->info.id_buku << endl;
                    cout << "Judul               : " << R->nextBuku->info.judul << endl;
                    cout << "Penulis             : " << R->nextBuku->info.penulis << endl;
                    cout << "Harga Sewa Per Hari : Rp " << R->nextBuku->info.harga_sewa << endl;
                }
                R = R->nextRelasi;
            }
            P = P->next;
            if (visitedPenyewa){
                cout << endl << "============================================================================================" << endl;
                cout << endl << "ENTER TO NEXT... ";
                getch();
                system("cls");
            }
            visitedPenyewa = false;
        }
    }
}

void showwAllBukuWithPenyewa(ListRelasi LR, ListBuku LB){
    /* I.S. Terdefinisi List Relasi (LR) dan ListBuku (LB)
       F.S. Semua data buku yang berelasi dengan penyewa ditampilkan ke layar beserta penyewanya.
            Jika tidak ada relasi, data buku tidak ditampilkan.
    */

    adrBuku B;
    adrRelasi R;
    bool visitedBuku = false;

    system("cls");
    if (LR.First == NULL){
        cout << endl << "============================================================================================" << endl;
        cout << "|*|                                 LIST RELASI KOSONG!!                                 |*|" << endl;
        cout << "============================================================================================" << endl;
        cout << endl << "ENTER TO NEXT... ";
        getch();
        system("cls");
    } else {
        B = LB.First;
        while (B != NULL){
            R = LR.First;
            while (R != NULL){
                if (R->nextBuku == B){
                    if (!visitedBuku){
                        cout << endl << "============================================================================================" << endl;
                        cout << "|*|                                      DATA BUKU                                       |*|" << endl;
                        cout << "============================================================================================" << endl;
                        cout << endl << "ID                  : " << B->info.id_buku << endl;
                        cout << "Judul               : " << B->info.judul << endl;
                        cout << "Penulis             : " << B->info.penulis << endl;
                        cout << "Harga Sewa Per Hari : Rp " << B->info.harga_sewa << endl;
                        cout << endl << "============================================================================================" << endl;

                        visitedBuku = true;

                        cout << endl << "============================================================================================" << endl;
                        cout << "|*|                            DATA YANG MENYEWA BUKU TERSEBUT                           |*|" << endl;
                        cout << "============================================================================================" << endl;
                    }
                    cout << endl << "ID      : " << R->nextPenyewa->info.id_penyewa << endl;
                    cout << "Nama    : " << R->nextPenyewa->info.nama << endl;
                    cout << "NIK     : " << R->nextPenyewa->info.nik << endl;
                }
                R = R->nextRelasi;
            }
            B = B->next;
            if (visitedBuku){
                cout << endl << "============================================================================================" << endl;
                cout << endl << "ENTER TO NEXT... ";
                getch();
                system("cls");
            }
            visitedBuku = false;
        }
    }
}

void countBukuDariPenyewa(ListRelasi LR, ListPenyewa LP, string id_penyewa){
    /* I.S. Terdefinisi List Relasi (LR) dan ListPenyewa (LP) yang mungkin kosong, serta string id_penyewa.
       F.S. Jumlah buku yang berelasi dengan penyewa tertentu ditampilkan di layar beserta seluruh infonya.
    */

    adrRelasi R;
    adrPenyewa P;
    int jumlahBuku  = 0;

    system("cls");
    if (LR.First == NULL){
        cout << endl << "============================================================================================" << endl;
        cout << "|*|                                 LIST RELASI KOSONG!!                                 |*|" << endl;
        cout << "============================================================================================" << endl;
    } else {
        P = findPenyewa(LP, id_penyewa);
        if (P != NULL){
            R = LR.First;
            while (R != NULL){
                if (R->nextPenyewa == P){
                    jumlahBuku++;
                }
                R = R->nextRelasi;
            }
            showBukuFromPenyewa(LR, LP, id_penyewa);
            cout << endl << "Jumlah Buku yang Disewa " << P->info.nama << " : " << jumlahBuku << endl;
            cout << endl << "============================================================================================" << endl;
        } else {
            cout << endl << "============================================================================================" << endl;
            cout << "|*|                             DATA PENYEWA TIDAK DITEMUKAN!                            |*|" << endl;
            cout << "============================================================================================" << endl;
        }
    }
    cout << endl << "ENTER TO NEXT... ";
    getch();
    system("cls");
}

void countPenyewaDariBuku(ListRelasi LR, ListBuku LB, string id_buku){
    /* I.S. Terdefinisi List Relasi (LR) yang mungkin kosong dan string id_buku.
       F.S. Jumlah penyewa yang berelasi dengan buku tertentu ditampilkan ke layar beserta seluruh infonya.
    */

    adrRelasi R = LR.First;
    adrBuku B;
    int hitung = 0;

    system("cls");
    if (LR.First == NULL){
        cout << endl << "============================================================================================" << endl;
        cout << "|*|                                 LIST RELASI KOSONG!!                                 |*|" << endl;
        cout << "============================================================================================" << endl;
    } else {
        B = findBuku(LB, id_buku);
        if (B != NULL){
            while (R != NULL){
                if (R ->nextBuku == B){
                    hitung++;
                }
                R = R -> nextRelasi;
            }
            showPenyewaFromBuku(LR, LB, id_buku);
            cout << endl << "Jumlah Penyewa Buku " << B->info.judul << " : " << hitung << endl;
            cout << endl << "============================================================================================" << endl;
        } else {
            cout << endl << "============================================================================================" << endl;
            cout << "|*|                              DATA BUKU TIDAK DITEMUKAN!                              |*|" << endl;
            cout << "============================================================================================" << endl;
        }
    }
    cout << endl << "ENTER TO NEXT... ";
    getch();
    system("cls");
}

void countBukuWithNoPenyewa(ListRelasi LR, ListBuku LB){
    /* I.S. Terdefinisi List Buku (LB) dan List Relasi (LR) yang mungkin kosong.
       F.S. Jumlah buku yang tidak memiliki relasi dengan penyewa ditampilkan ke layar beserta seluruh infonya.
    */

    adrBuku B;
    adrRelasi R;
    int jumlahNoPenyewa = 0;
    bool adaPenyewa = false;

    system("cls");
    if (LR.First == NULL){
        cout << endl << "============================================================================================" << endl;
        cout << "|*|                                 LIST RELASI KOSONG!!                                 |*|" << endl;
        cout << "============================================================================================" << endl;
    } else {
        B = LB.First;
        while (B != NULL){
            R = LR.First;
            while (R != NULL && !adaPenyewa){
                adaPenyewa = R->nextBuku == B && R->nextPenyewa != NULL;
                if (!adaPenyewa){
                    R = R->nextRelasi;
                }
            }
            if (!adaPenyewa){
                jumlahNoPenyewa++;
                if (jumlahNoPenyewa == 1){
                    cout << endl << "============================================================================================" << endl;
                    cout << "|*|                               DATA BUKU TANPA PENYEWA                                |*|" << endl;
                    cout << "============================================================================================" << endl;
                }
                cout << endl << "ID                  : " << B->info.id_buku << endl;
                cout << "Judul               : " << B->info.judul << endl;
                cout << "Penulis             : " << B->info.penulis << endl;
                cout << "Harga Sewa Per Hari : Rp " << B->info.harga_sewa << endl;
            }
            B = B->next;
            adaPenyewa = false;
        }

        if (jumlahNoPenyewa == 0){
            cout << endl << "============================================================================================" << endl;
            cout << "|*|                             TIDAK ADA BUKU TANPA PENYEWA!                            |*|" << endl;
            cout << "============================================================================================" << endl;
        } else {
            cout << endl << "============================================================================================" << endl;
        }

        cout << endl << "Jumlah Buku Tanpa Penyewa : " << jumlahNoPenyewa << endl;
        cout << endl << "============================================================================================" << endl;
    }
    cout << endl << "ENTER TO NEXT... ";
    getch();
    system("cls");
}

void countPenyewaWithNoBuku(ListRelasi LR, ListPenyewa LP){
    /* I.S. Terdefinisi List Penyewa (LB) dan List Relasi (LR) yang mungkin kosong.
       F.S. Jumlah penyewa yang tidak memiliki relasi dengan buku ditampilkan ke layar beserta seluruh infonya.
    */

    adrPenyewa P;
    adrRelasi R;
    int jumlahNoBuku = 0;
    bool adaBuku = false;

    system("cls");
    if (LR.First == NULL){
        cout << endl << "============================================================================================" << endl;
        cout << "|*|                                 LIST RELASI KOSONG!!                                 |*|" << endl;
        cout << "============================================================================================" << endl;
    } else {
        P = LP.First;
        while (P != NULL){
            R = LR.First;
            while (R != NULL && !adaBuku){
                adaBuku = R->nextPenyewa == P;
                if (!adaBuku){
                    R = R->nextRelasi;
                }
            }
            if (!adaBuku){
                jumlahNoBuku++;
                if (jumlahNoBuku == 1){
                    cout << endl << "============================================================================================" << endl;
                    cout << "|*|                         DATA PENYEWA YANG TIDAK MENYEWA BUKU!                        |*|" << endl;
                    cout << "============================================================================================" << endl;
                }
                cout << endl << "ID      : " << P->info.id_penyewa << endl;
                cout << "Nama    : " << P->info.nama << endl;
                cout << "NIK     : " << P->info.nik << endl;
            }
            P = P->next;
            adaBuku = false;
        }

        if (jumlahNoBuku == 0){
            cout << endl << "============================================================================================" << endl;
            cout << "|*|                      TIDAK ADA PENYEWA YANG TIDAK MENYEWA BUKU!                      |*|" << endl;
            cout << "============================================================================================" << endl;
        } else {
            cout << endl << "============================================================================================" << endl;
        }

        cout << endl << "Jumlah Penyewa yang Tidak Menyewa Buku : " << jumlahNoBuku << endl;
        cout << endl << "============================================================================================" << endl;
    }
    cout << endl << "ENTER TO NEXT... ";
    getch();
    system("cls");
}

void editBukuDariPenyewa(ListRelasi &LR, ListBuku LB, string id_penyewa, string id_buku, string newId_Buku){
    /* I.S. List LR dan LB mungkin kosong, dengan id_penyewa, id_buku, dan newId_Buku terdefinisi.
       F.S. Relasi id_penyewa dengan id_buku dipindahkan ke id_penyewa dengan newId_Buku, jika relasi dan buku baru ditemukan.
            Jika tidak ditemukan, ditampilkan pesan sesuai kondisi.
    */

    adrBuku B;
    adrRelasi R;
    R = findRelasiPenyewaAndBuku(LR, id_penyewa, id_buku);

    system("cls");
    if (LR.First == NULL){
        cout << endl << "============================================================================================" << endl;
        cout << "|*|                                 LIST RELASI KOSONG!!                                 |*|" << endl;
        cout << "============================================================================================" << endl;
    } else {
        if (R == NULL){
            cout << endl << "============================================================================================" << endl;
            cout << "|*|                             DATA RELASI TIDAK DITEMUKAN!                             |*|" << endl;
            cout << "============================================================================================" << endl;
        } else {
            B = findBuku(LB, newId_Buku);
            if (B == NULL){
                cout << endl << "============================================================================================" << endl;
                cout << "|*|                 DATA RELASI DITEMUKAN, TETAPI DATA BUKU BARU TIDAK!!                 |*|" << endl;
                cout << "============================================================================================" << endl;
            } else {
                cout << endl << "============================================================================================" << endl;
                cout << "|*|                          DATA RELASI DAN BUKU BARU DITEMUKAN!                        |*|" << endl;
                cout << "============================================================================================" << endl;
                cout << endl << "ENTER TO NEXT... ";
                getch();
                system("cls");

                R ->nextBuku = B;

                cout << endl << "============================================================================================" << endl;
                cout << "|*|                             DATA RELASI BERHASIL DIEDIT!                             |*|" << endl;
                cout << "============================================================================================" << endl;
            }
        }
    }
    cout << endl << "ENTER TO NEXT... ";
    getch();
    system("cls");

}

void editPenyewaDariBuku(ListRelasi &LR, ListPenyewa LP, string id_buku, string id_penyewa, string newId_Penyewa){
    /* I.S. List LR dan LP mungkin kosong, dengan id_buku, id_penyewa, dan newId_Penyewa terdefinisi.
       F.S. Relasi id_buku dengan id_penyewa dipindahkan ke id_buku dengan newId_Penyewa, jika relasi dan penyewa baru ditemukan.
            Jika tidak ditemukan, ditampilkan pesan sesuai kondisi.
    */

    adrPenyewa P;
    adrRelasi R;

    system("cls");
    if (LR.First == NULL){
        cout << endl << "============================================================================================" << endl;
        cout << "|*|                                 LIST RELASI KOSONG!!                                 |*|" << endl;
        cout << "============================================================================================" << endl;
    } else {
        R = findRelasiPenyewaAndBuku(LR, id_penyewa, id_buku);
        if (R == NULL){
            cout << endl << "============================================================================================" << endl;
            cout << "|*|                             DATA RELASI TIDAK DITEMUKAN!                             |*|" << endl;
            cout << "============================================================================================" << endl;
        } else {
            P = findPenyewa(LP, newId_Penyewa);
            if (P == NULL){
                cout << endl << "============================================================================================" << endl;
                cout << "|*|                DATA RELASI DITEMUKAN, TETAPI DATA PENYEWA BARU TIDAK!                |*|" << endl;
                cout << "============================================================================================" << endl;
            } else {
                cout << endl << "============================================================================================" << endl;
                cout << "|*|                        DATA RELASI DAN PENYEWA BARU DITEMUKAN!                       |*|" << endl;
                cout << "============================================================================================" << endl;
                cout << endl << "ENTER TO NEXT... ";
                getch();
                system("cls");

                R ->nextPenyewa = P;

                cout << endl << "============================================================================================" << endl;
                cout << "|*|                             DATA RELASI BERHASIL DIEDIT!                             |*|" << endl;
                cout << "============================================================================================" << endl;
            }
        }
    }
    cout << endl << "ENTER TO NEXT... ";
    getch();
    system("cls");
}

void insertLastRelasi(ListRelasi &LR, adrRelasi R){
    /* I.S. Terdefinisi List Relasi (LR) yang mungkin kosong dan sebuah elemen baru.
       F.S. Elemen relasi baru telah ditambahkan sebagai elemen terakhir dalam LR.
    */
    adrRelasi Q;
    if (LR.First == NULL){
        LR.First = R;
    } else {
        Q = LR.First;
        while (Q->nextRelasi != NULL){
            Q = Q->nextRelasi;
        }
        Q->nextRelasi = R;
    }
}

void deleteFirstPenyewa(ListPenyewa &LP, adrPenyewa &P){
    /*I.S. Terdefinisi List LP yang mungkin kosong atau hanya memiliki 1 elemen
      F.S. Jika list kosong, maka pointer P di assign dengan NIL, jika tidak kosong
           maka elemen paling awal dihapus dari List LP, alamatnya disimpan oleh pointer P
    */
     if (LP.First == NULL){
        P = NULL;
    }else if (LP.First->next == NULL){
        P = LP.First;
        LP.First = NULL;
    }else{
        P = LP.First;
        LP.First = P->next;
        P->next = NULL;
    }

}

void deleteLastPenyewa(ListPenyewa &LP, adrPenyewa &P){
     /*I.S. Terdefinisi List LP yang mungkin kosong atau hanya memiliki 1 elemen
       F.S. Jika list kosong, maka pointer P di assign dengan NIL, jika tidak kosong
            maka elemen paling akhir dihapus dari List LP, alamatnya disimpan oleh pointer P
    */
     adrPenyewa Q, S;
    if (LP.First == NULL){
        P = NULL;
    }else if (LP.First->next == NULL){
        P = LP.First;
        LP.First = NULL;
    }else {
        Q = LP.First;
        while (Q->next != NULL){
            S = Q;
            Q = Q->next;
        }
        P = S->next;
        S->next = NULL;
    }

}

void deleteAfterPenyewa(ListPenyewa &LP, adrPenyewa prec, adrPenyewa &P){
     /*I.S. Terdefinisi List LP yang mungkin kosong, sebuah elemen baru yang alamatnya disimpan
            oleh pointer P dan pointer prec
       F.S. Elemen telah dihapus dari list LP, alamatnya disimpan oleh pointer P
    */
    if (prec == NULL || prec->next == NULL){
        P = NULL;
    }else {
        P  = prec->next;
        prec->next = P->next;
        P->next = NULL;
    }

}

void deleteFirstBuku(ListBuku &LB, adrBuku &B){
    /*I.S. Terdefinisi List LB yang mungkin kosong atau hanya memiliki 1 elemen
      F.S. Jika list kosong, maka pointer B di assign dengan NIL, jika tidak kosong
           maka elemen paling awal dihapus dari List LB, alamatnya disimpan oleh pointer B
    */
    if (LB.First == NULL){
        B = NULL;
    }else if (LB.First == LB.Last){
        B = LB.First;
        LB.First = NULL;
        LB.Last = NULL;
    } else{
        B = LB.First;
        LB.First = LB.First -> next;
        LB.First -> prev = NULL;
        B -> next = NULL;
    }
}

void deleteLastBuku(ListBuku &LB, adrBuku &B){
     /*I.S. Terdefinisi List LB yang mungkin kosong atau hanya memiliki 1 elemen
       F.S. Jika list kosong, maka pointer B di assign dengan NIL, jika tidak kosong
            maka elemen paling akhir dihapus dari List LB, alamatnya disimpan oleh pointer B
    */
    if (LB.First == NULL){
        B = NULL;
    }else if (LB.First == LB.Last){
        B = LB.First;
        LB.First = NULL;
        LB.Last = NULL;
    } else{
        B = LB.Last;
        LB.Last = LB.Last -> prev;
        LB.Last -> next = NULL;
        B -> prev = NULL;
    }
}

void deleteAfterBuku(ListBuku &LB, adrBuku prec, adrBuku &B){
    /*I.S. Terdefinisi List LB yang mungkin kosong, sebuah elemen baru yang alamatnya disimpan
           oleh pointer B dan pointer prec
      F.S. Elemen telah dihapus dari list LB, alamatnya disimpan oleh pointer B
    */
    if (prec == NULL || prec->next == NULL){
        B = NULL;
    } else {
        B = prec->next;
        prec->next = B->next;
        if (B->next != NULL) {
            B->next->prev = prec;
        } else {
            LB.Last = prec;
        }
        B->next = NULL;
        B->prev = NULL;
    }
}

void deleteFirstRelasi(ListRelasi &LR, adrRelasi &R){
    /*I.S. Terdefinisi List LR yang mungkin kosong atau hanya memiliki 1 elemen
      F.S. Jika list kosong, maka pointer R di assign dengan NIL, jika tidak kosong
           maka elemen paling awal dihapus dari List LR, alamatnya disimpan oleh pointer R
    */
    if (LR.First == NULL){
        R = NULL;
    } else if (LR.First->nextRelasi == NULL){
        R = LR.First;
        R->nextBuku = NULL;
        R->nextPenyewa = NULL;
        LR.First = NULL;
    } else {
        R = LR.First;
        LR.First = R->nextRelasi;
        R->nextBuku = NULL;
        R->nextPenyewa = NULL;
        R->nextRelasi = NULL;
    }
}

void deleteLastRelasi(ListRelasi &LR, adrRelasi &R){
     /*I.S. Terdefinisi List LR yang mungkin kosong atau hanya memiliki 1 elemen
       F.S. Jika list kosong, maka pointer R di assign dengan NIL, jika tidak kosong
            maka elemen paling akhir dihapus dari List LR, alamatnya disimpan oleh pointer R
    */
    adrRelasi Q, S;
    if (LR.First == NULL){
        R = NULL;
    } else if (LR.First->nextRelasi == NULL){
        R = LR.First;
        R->nextBuku = NULL;
        R->nextPenyewa = NULL;
        LR.First = NULL;
    } else {
        Q = LR.First;
        while (Q->nextRelasi != NULL){
            S = Q;
            Q = Q->nextRelasi;
        }
        R = S->nextRelasi;
        R->nextBuku = NULL;
        R->nextPenyewa = NULL;
        S->nextRelasi = NULL;
    }
}

void deleteAfterRelasi(ListRelasi &LR, adrRelasi prec, adrRelasi &R){
     /*I.S. Terdefinisi List LR yang mungkin kosong, sebuah elemen baru yang alamatnya disimpan
            oleh pointer R dan pointer prec
       F.S. Elemen telah dihapus dari list LR, alamatnya disimpan oleh pointer R
    */
    if (prec == NULL || prec->nextRelasi == NULL){
        R = NULL;
    } else {
        R = prec->nextRelasi;
        prec->nextRelasi = R->nextRelasi;
        R->nextBuku = NULL;
        R->nextPenyewa = NULL;
        R->nextRelasi = NULL;
    }
}

void showOneDataPenyewa(adrPenyewa P){
    /* I.S. Terdefinisi pointer P yang mungkin NULL atau menunjuk ke elemen data penyewa.
       F.S. Menampilkan data penyewa jika P tidak NULL, atau pesan "DATA PENYEWA TIDAK DITEMUKAN" jika P NULL.
    */

    if (P == NULL){
        cout << endl << "============================================================================================" << endl;
        cout << "|*|                             DATA PENYEWA TIDAK DITEMUKAN!                            |*|" << endl;
        cout << "============================================================================================" << endl;
    } else {
        cout << endl << "============================================================================================" << endl;
        cout << "|*|                                     DATA PENYEWA                                     |*|" << endl;
        cout << "============================================================================================" << endl;
        cout << endl << "ID      : " << P->info.id_penyewa << endl;
        cout << "Nama    : " << P->info.nama << endl;
        cout << "NIK     : " << P->info.nik << endl;
        cout << endl << "============================================================================================" << endl;
    }
}

void showOneDataBuku(adrBuku B){
    /* I.S. Terdefinisi pointer B yang mungkin NULL atau menunjuk ke elemen data buku.
       F.S. Menampilkan data buku jika B tidak NULL, atau pesan "DATA BUKU TIDAK DITEMUKAN" jika B NULL.
    */

    if (B == NULL){
        cout << endl << "============================================================================================" << endl;
        cout << "|*|                              DATA BUKU TIDAK DITEMUKAN!                              |*|" << endl;
        cout << "============================================================================================" << endl;
    } else {
        cout << endl << "============================================================================================" << endl;
        cout << "|*|                                      DATA BUKU                                       |*|" << endl;
        cout << "============================================================================================" << endl;
        cout << endl << "ID                  : " << B->info.id_buku << endl;
        cout << "Judul               : " << B->info.judul << endl;
        cout << "Penulis             : " << B->info.penulis << endl;
        cout << "Harga Sewa Per Hari : Rp " << B->info.harga_sewa << endl;
        cout << endl << "============================================================================================" << endl;
    }
}

void showOneDataRelasi(adrRelasi R){
    /* I.S. Terdefinisi pointer R yang mungkin NULL atau menunjuk ke elemen data relasi.
       F.S. Menampilkan data penyewa dan buku dari relasi jika R tidak NULL, atau pesan "DATA RELASI TIDAK DITEMUKAN" jika R NULL.
    */

    system("cls");
    if (R == NULL){
        cout << endl << "============================================================================================" << endl;
        cout << "|*|                             DATA RELASI TIDAK DITEMUKAN!                             |*|" << endl;
        cout << "============================================================================================" << endl;
    } else {
        cout << endl << "============================================================================================" << endl;
        cout << "|*|                                DATA RELASI DITEMUKAN!                                |*|" << endl;
        cout << "============================================================================================" << endl;

        cout << endl << "ENTER TO NEXT... ";
        getch();
        system("cls");

        cout << endl << "============================================================================================" << endl;
        cout << "|*|                                     DATA PENYEWA                                     |*|" << endl;
        cout << "============================================================================================" << endl;
        cout << endl << "ID      : " << R->nextPenyewa->info.id_penyewa << endl;
        cout << "Nama    : " << R->nextPenyewa->info.nama << endl;
        cout << "NIK     : " << R->nextPenyewa->info.nik << endl;
        cout << endl << "============================================================================================" << endl;

        cout << endl << "============================================================================================" << endl;
        cout << "|*|                                      DATA BUKU                                       |*|" << endl;
        cout << "============================================================================================" << endl;
        cout << endl << "ID                  : " << R->nextBuku->info.id_buku << endl;
        cout << "Judul               : " << R->nextBuku->info.judul << endl;
        cout << "Penulis             : " << R->nextBuku->info.penulis << endl;
        cout << "Harga Sewa Per Hari : Rp " << R->nextBuku->info.harga_sewa << endl;
        cout << endl << "============================================================================================" << endl;
    }
    cout << endl << "ENTER TO NEXT... ";
    getch();
    system("cls");
}

void validIdPenyewa(ListPenyewa LP, string &id_penyewa){
    /* I.S. List LP terdefinisi, id_penyewa merupakan input awal pengguna.
       F.S. id_penyewa menjadi valid (tidak duplikat) berdasarkan data dalam ListPenyewa (LP).
    */

    adrPenyewa P;
    P = findPenyewa(LP, id_penyewa);
    while (P != NULL){
        system("cls");
        cout << endl << "============================================================================================" << endl;
        cout << "|*|                                ID PENYEWA TIDAK VALID!                               |*|" << endl;
        cout << "|*|                                 SILAHKAN INPUT ULANG!                                |*|" << endl;
        cout << "============================================================================================" << endl;

        cout << endl << "ENTER TO NEXT... ";
        getch();
        system("cls");

        cout << endl << "============================================================================================" << endl;
        cout << "|*|                                  INPUT DATA PENYEWA                                  |*|" << endl;
        cout << "============================================================================================" << endl;


        cout << endl << "Masukkan ID Penyewa : ";
        cin >> id_penyewa;
        P = findPenyewa(LP, id_penyewa);
    }
}

void validIdBuku(ListBuku LB, string &id_buku){
    /* I.S. List LB terdefinisi, id_buku merupakan input awal pengguna.
       F.S. id_buku menjadi valid (tidak duplikat) berdasarkan data dalam ListBuku (LB).
    */
    adrBuku B;
    B = findBuku(LB, id_buku);
    while (B != NULL){
        system("cls");
        cout << endl << "============================================================================================" << endl;
        cout << "|*|                                 ID BUKU TIDAK VALID!                                 |*|" << endl;
        cout << "|*|                                 SILAHKAN INPUT ULANG!                                |*|" << endl;
        cout << "============================================================================================" << endl;

        cout << endl << "ENTER TO NEXT... ";
        getch();
        system("cls");

        cout << endl << "============================================================================================" << endl;
        cout << "|*|                                  TAMBAH DATA BUKU                                    |*|" << endl;
        cout << "============================================================================================" << endl;

        cout << endl << "Masukkan ID Buku : ";
        cin >> id_buku;
        B = findBuku(LB, id_buku);
    }
}

int menuUtama(){
    /* I.S. -
       F.S. Menampilkan menu utama dan mengembalikan pilihan pengguna dalam rentang [0-3].
    */

    system("cls");
    cout << "============================================================================================" << endl;
    cout << "|*|                                                                                      |*|" << endl;
    cout << "|*|                                      MENU UTAMA                                      |*|" << endl;
    cout << "|*|                                                                                      |*|" << endl;
    cout << "============================================================================================" << endl;
    cout << "|*|                                                                                      |*|" << endl;
    cout << "|*|  1. Menu Penyewa                                                                     |*|" << endl;
    cout << "|*|  2. Menu Buku                                                                        |*|" << endl;
    cout << "|*|  3. Menu Relasi                                                                      |*|" << endl;
    cout << "|*|  0. Exit Program                                                                     |*|" << endl;
    cout << "|*|                                                                                      |*|" << endl;
    cout << "============================================================================================" << endl;

    int input = 0;
    cout << "Masukkan Pilihan (1/../3/0) : ";
    cin >> input;
    while (input < 0 || input > 3){
        cout << endl << "============================================================================================" << endl;
        cout << "|*|                                 PILIHAN TIDAK VALID!                                 |*|" << endl;
        cout << "|*|                                SILAHKAN MEMILIH ULANG!                               |*|" << endl;
        cout << "============================================================================================" << endl;
        cout << endl << "Masukkan Pilihan (1/../3/0) : ";
        cin >> input;
    }

    return input;

}

int menuPenyewa(){
    /* I.S. -
       F.S. Menampilkan menu penyewa dan mengembalikan pilihan pengguna dalam rentang [0-4].
    */

    cout << "============================================================================================" << endl;
    cout << "|*|                                                                                      |*|" << endl;
    cout << "|*|                                     MENU PENYEWA                                     |*|" << endl;
    cout << "|*|                                                                                      |*|" << endl;
    cout << "============================================================================================" << endl;
    cout << "|*|                                                                                      |*|" << endl;
    cout << "|*|  1. Tambah Data Penyewa                                                              |*|" << endl;
    cout << "|*|  2. Hapus Data Penyewa                                                               |*|" << endl;
    cout << "|*|  3. Cari Data Penyewa                                                                |*|" << endl;
    cout << "|*|  4. Lihat Data Penyewa                                                               |*|" << endl;
    cout << "|*|  0. Exit Menu                                                                        |*|" << endl;
    cout << "|*|                                                                                      |*|" << endl;
    cout << "============================================================================================" << endl;

    int input = 0;
    cout << "Masukkan Pilihan (1/../4/0) : ";
    cin >> input;
    while (input < 0 || input > 4){
        cout << endl << "============================================================================================" << endl;
        cout << "|*|                                 PILIHAN TIDAK VALID!                                 |*|" << endl;
        cout << "|*|                                SILAHKAN MEMILIH ULANG!                               |*|" << endl;
        cout << "============================================================================================" << endl;
        cout << endl << "Masukkan Pilihan (1/../4/0) : ";
        cin >> input;
    }
    return input;
}

int menuBuku(){
    /* I.S. -
       F.S. Menampilkan menu buku dan mengembalikan pilihan pengguna dalam rentang [0-4].
    */

    cout << "============================================================================================" << endl;
    cout << "|*|                                                                                      |*|" << endl;
    cout << "|*|                                      MENU  BUKU                                      |*|" << endl;
    cout << "|*|                                                                                      |*|" << endl;
    cout << "============================================================================================" << endl;
    cout << "|*|                                                                                      |*|" << endl;
    cout << "|*|  1. Tambah Data Buku                                                                 |*|" << endl;
    cout << "|*|  2. Hapus Data Buku                                                                  |*|" << endl;
    cout << "|*|  3. Cari Data Buku                                                                   |*|" << endl;
    cout << "|*|  4. Lihat Data Buku                                                                  |*|" << endl;
    cout << "|*|  0. Exit Menu                                                                        |*|" << endl;
    cout << "|*|                                                                                      |*|" << endl;
    cout << "============================================================================================" << endl;

    int input = 0;
    cout << "Masukkan Pilihan (1/../4/0) : ";
    cin >> input;
    while (input < 0 || input > 4){
        cout << endl << "============================================================================================" << endl;
        cout << "|*|                                 PILIHAN TIDAK VALID!                                 |*|" << endl;
        cout << "|*|                                SILAHKAN MEMILIH ULANG!                               |*|" << endl;
        cout << "============================================================================================" << endl;
        cout << endl << "Masukkan Pilihan (1/../4/0) : ";
        cin >> input;
    }
    return input;
}

int menuRelasi(){
    /* I.S. -
       F.S. Menampilkan menu relasi dan mengembalikan pilihan pengguna dalam rentang [0-13].
    */

    cout << "============================================================================================" << endl;
    cout << "|*|                                                                                      |*|" << endl;
    cout << "|*|                                     MENU  RELASI                                     |*|" << endl;
    cout << "|*|                                                                                      |*|" << endl;
    cout << "============================================================================================" << endl;
    cout << "|*|                                                                                      |*|" << endl;
    cout << "|*|  1.  Tambah Relasi Antara Penyewa dan Buku Tertentu                                  |*|" << endl;
    cout << "|*|  2.  Hapus Relasi Antara Penyewa dan Buku Tertentu                                   |*|" << endl;
    cout << "|*|  3.  Cari Relasi Antara Penyewa dan Buku Tertentu                                    |*|" << endl;
    cout << "|*|  4.  Lihat Data Buku dari Penyewa Tertentu                                           |*|" << endl;
    cout << "|*|  5.  Lihat Data Penyewa dari Buku Tertentu                                           |*|" << endl;
    cout << "|*|  6.  Lihat Semua Data Penyewa dengan Buku yang Disewanya                             |*|" << endl;
    cout << "|*|  7.  Lihat Semua Data Buku dengan Penyewanya                                         |*|" << endl;
    cout << "|*|  8.  Lihat Jumlah Buku dari Penyewa Tertentu                                         |*|" << endl;
    cout << "|*|  9.  Lihat Jumlah Penyewa dari Buku Tertentu                                         |*|" << endl;
    cout << "|*|  10. Lihat Jumlah Buku yang Tidak Disewa                                             |*|" << endl;
    cout << "|*|  11. Lihat Jumlah Penyewa yang Tidak Menyewa Buku                                    |*|" << endl;
    cout << "|*|  12. Edit Data Buku dari Penyewa Tertentu                                            |*|" << endl;
    cout << "|*|  13. Edit Data Penyewa dari Buku Tertentu                                            |*|" << endl;
    cout << "|*|  0.  Exit Menu                                                                       |*|" << endl;
    cout << "|*|                                                                                      |*|" << endl;
    cout << "============================================================================================" << endl;

    int input = 0;
    cout << "Masukkan Pilihan (1/../13/0) : ";
    cin >> input;
    while (input < 0 || input > 13){
        cout << endl << "============================================================================================" << endl;
        cout << "|*|                                 PILIHAN TIDAK VALID!                                 |*|" << endl;
        cout << "|*|                                SILAHKAN MEMILIH ULANG!                               |*|" << endl;
        cout << "============================================================================================" << endl;
        cout << endl << "Masukkan Pilihan (1/../13/0) : ";
        cin >> input;
    }
    return input;
}
