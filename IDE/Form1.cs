using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;
using System.IO;
using System.Collections;
using System.Text.RegularExpressions;

namespace IDE
{
    public partial class Form1 : Form
    {
        private ArrayList Palabra = new ArrayList();
        int indice = 0;
        private Dictionary<string, Color> PalabrasReservadas = new Dictionary<string, Color>();
        public Form1()
        {
            InitializeComponent();
        }

        private void archivoToolStripMenuItem_Click(object sender, EventArgs e)
        {

        }

        private void ScannerDoc()
        {
            foreach (KeyValuePair<string, Color> PalReserv in PalabrasReservadas)
            {
                MatchCollection Resultados;
                Regex busqueda = new Regex(PalReserv.Key.ToString(), RegexOptions.IgnoreCase);
                Resultados = busqueda.Matches(richTextBox1.Text);
                richTextBox1.SelectAll();

                foreach (Match Palabra in Resultados)
                {
                    richTextBox1.SelectionStart = Palabra.Index;
                    richTextBox1.SelectionLength = Palabra.Length;

                    richTextBox1.SelectionColor = PalReserv.Value;

                    richTextBox1.SelectionStart = Palabra.Length + 100;
                    richTextBox1.SelectionLength = Palabra.Length + 101;

                }
            }

        }

        private void Form1_Load(object sender, EventArgs e)
        {
            PalabrasReservadas.Add("int", Color.Blue);
            PalabrasReservadas.Add("String", Color.LightBlue);
            PalabrasReservadas.Add("for", Color.DarkViolet);
            PalabrasReservadas.Add("Float", Color.ForestGreen);
            PalabrasReservadas.Add("Double", Color.ForestGreen);
            PalabrasReservadas.Add("if", Color.DarkViolet);
            PalabrasReservadas.Add("else", Color.DarkViolet);
            PalabrasReservadas.Add("while", Color.DarkViolet);
            PalabrasReservadas.Add("switch", Color.DarkViolet);
            PalabrasReservadas.Add("case", Color.Blue);
            PalabrasReservadas.Add("cout", Color.Red);

        }

        private void btn2_Click(object sender, EventArgs e)
        {
            MessageBox.Show("Hola Mundo");
        }

        private void nuevoToolStripMenuItem_Click(object sender, EventArgs e)
        {
            
        }

        private void abrirToolStripMenuItem1_Click(object sender, EventArgs e)
        {
            //Salvamos el Documento text
            richTextBox1.SaveFile(openFileDialog1.FileName, RichTextBoxStreamType.RichText);

        }

        

        private void toolStripButton7_Click(object sender, EventArgs e)
        {
            

        }

        private void abrirToolStripMenuItem_Click(object sender, EventArgs e)
        {
            if (openFileDialog1.ShowDialog() == DialogResult.OK)
            {
                //Cargamos el documento rtf
                richTextBox1.LoadFile(openFileDialog1.FileName);
                ScannerDoc();
                
            }

        }

        private void toolStripButton1_Click(object sender, EventArgs e)
        {
            
        }

        private void toolStripButton3_Click(object sender, EventArgs e)
        {
            if (openFileDialog1.ShowDialog() == DialogResult.OK)
            {
                //Cargamos el documento rtf
                richTextBox1.LoadFile(openFileDialog1.FileName);
            }
        }

        private void fuenteToolStripMenuItem_Click(object sender, EventArgs e)
        {
            if(fontDialog1.ShowDialog() == DialogResult.OK)
            {
                richTextBox1.SelectionFont = fontDialog1.Font;
            }
        }

        private void colorToolStripMenuItem_Click(object sender, EventArgs e)
        {
            if(colorDialog1.ShowDialog()== DialogResult.OK)
            {
                richTextBox1.SelectionColor = colorDialog1.Color;


            }
        }

        private void toolStrip1_ItemClicked(object sender, ToolStripItemClickedEventArgs e)
        {

        }

        private void toolStripButton2_Click(object sender, EventArgs e)
        {
            SaveFileDialog guardar = new SaveFileDialog();


            guardar.Filter = "Documento de texto|*.txt";
            guardar.Title = "Guardar RichTextBox";
            guardar.FileName = "Sin titulo 1";
            var resultado = guardar.ShowDialog();
            if (resultado == DialogResult.OK)
            {

                StreamWriter escribir = new StreamWriter(guardar.FileName);
                foreach (object line in richTextBox1.Lines)
                {
                    escribir.WriteLine(line);
                }

                escribir.Close();
            }

        }

        private void toolStripButton4_Click(object sender, EventArgs e)
        {
            OpenFileDialog abrir = new OpenFileDialog();
            abrir.Filter = "Documento de texto|*.txt";
            abrir.Title = "Abrir";
            abrir.FileName = "Sin titulo 1";
            var resultado = abrir.ShowDialog();

            if (resultado == DialogResult.OK)
            {
                StreamReader leer = new StreamReader(abrir.FileName);
                richTextBox1.Text = leer.ReadToEnd();
         
                leer.Close();
                ScannerDoc();
            }
        }

        private void salirToolStripMenuItem_Click(object sender, EventArgs e)
        {
            SaveFileDialog guardar = new SaveFileDialog();


            guardar.Filter = "Documento de texto|*.txt";
            guardar.Title = "Guardar RichTextBox";
            guardar.FileName = "Sin titulo 1";
            if (saveFileDialog1.ShowDialog() == DialogResult.OK)
            {
                StreamWriter escribir = new StreamWriter(guardar.FileName);
                foreach (object line in richTextBox1.Lines)
                {
                    escribir.WriteLine(line);
                }

                escribir.Close();

                //Salvamos el Documento text
                richTextBox1.SaveFile(saveFileDialog1.FileName, RichTextBoxStreamType.RichText);
                

            }
        }

        private void checaPalabra()
        {
            /*String Pal = "";
            Boolean band = false;
            for (int i = 0; i < Palabra.Count; i++)
            {
                Pal += (String)Palabra[i];
            }
            //MessageBox.Show(Pal);
            foreach (KeyValuePair<string, Color> PalReserv in PalabrasReservadas)
            {
                if (PalReserv.Key.Equals(Pal))
                {
                    //MessageBox.Show("Encontrado");
                    MatchCollection Resultados;
                    Regex busqueda = new Regex(PalReserv.Key.ToString(), RegexOptions.IgnoreCase);
                    Resultados = busqueda.Matches(richTextBox1.Text);

                    foreach (Match Palabra in Resultados)
                    {
                        richTextBox1.SelectionStart = Palabra.Index;
                        richTextBox1.SelectionLength = Palabra.Length;

                        richTextBox1.SelectionColor = PalReserv.Value;
                        //richTextBox1.SelectionStart = this.richTextBox1.Text.IndexOf(Pal);
                        //richTextBox1.SelectionStart = this.richTextBox1.Text.IndexOf(Pal);

                        richTextBox1.SelectionStart = Palabra.Length + 100;
                        richTextBox1.SelectionLength = Palabra.Length + 100;
                        richTextBox1.SelectionColor = Color.Black;

                    }
                    Palabra.Clear();
                    band = true;
                    break;
                }
                else
                {
                    richTextBox1.SelectionColor = Color.Black;
                }

            }
            Palabra.Clear();
        }*/
            string comparar = "";
            int indice = 0;
            comparar = richTextBox1.Text;
            char[] arreglo = richTextBox1.Text.ToCharArray();
            for (int i = 0; i < richTextBox1.TextLength; i++)
            {
                if (arreglo[i] == ' ')
                {
                    indice = i;
                }
                else
                {
                    if (arreglo[i] == '\n')
                    {
                        indice = i;
                    }
                }
            }

            comparar = null;
            if (indice != 0)
            {
                indice++;
            }
            for (int i = indice; i < richTextBox1.Text.Length; i++)
            {
                comparar += arreglo[i];
            }


            compararTexto(comparar, indice);

        }
        


        private void compararTexto(string comparar, int indice)
        {
          

           foreach (KeyValuePair<string, Color> PalReserv in PalabrasReservadas)
            {
                if (PalReserv.Key.Equals(comparar))
                {
                    richTextBox1.Select(indice, indice+comparar.Length);
                    richTextBox1.SelectionColor = PalReserv.Value;
                    richTextBox1.SelectionStart = this.richTextBox1.Text.Length;
                    richTextBox1.SelectionColor = Color.Black;
                }

                }
           }
   

        private void richTextBox1_KeyPress(object sender, KeyPressEventArgs e)
        {
            if(e.KeyChar == Convert.ToChar(Keys.Space))
            {
           checaPalabra();
            }
            else
            {
                if(e.KeyChar == Convert.ToChar(Keys.Enter))
                {
                    //indice = 0;
                    Palabra.Clear();
                }
                else
                {
                    if(e.KeyChar == Convert.ToChar(Keys.Back))
                    {
                        if (Palabra.Count > 0)
                        {
                            Palabra.RemoveAt(Palabra.Count - 1);
                        }
                        
                    }
                    else
                    {
                        Palabra.Add(e.KeyChar.ToString());
                       // checaPalabra();
                    }
                }
                
                
                
            }
        }

        
    }
}
