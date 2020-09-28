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
        private string rutaArchivo ="";
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
            PalabrasReservadas.Add("string", Color.LightBlue);
            PalabrasReservadas.Add("for", Color.DarkViolet);
            PalabrasReservadas.Add("float", Color.ForestGreen);
            PalabrasReservadas.Add("double", Color.ForestGreen);
            PalabrasReservadas.Add("if", Color.DarkViolet);
            PalabrasReservadas.Add("else", Color.DarkViolet);
            PalabrasReservadas.Add("while", Color.DarkViolet);
            PalabrasReservadas.Add("switch", Color.DarkViolet);
            PalabrasReservadas.Add("case", Color.Blue);
            PalabrasReservadas.Add("cout", Color.Red);
            PalabrasReservadas.Add("program", Color.Red);
            PalabrasReservadas.Add("bool", Color.Blue);
            PalabrasReservadas.Add("fi", Color.Blue);
            PalabrasReservadas.Add("until", Color.Blue);
            PalabrasReservadas.Add("read", Color.Blue);
            PalabrasReservadas.Add("write", Color.Blue);
            PalabrasReservadas.Add("not", Color.Blue);
            PalabrasReservadas.Add("and", Color.Blue);
            PalabrasReservadas.Add("then", Color.Blue);
            PalabrasReservadas.Add("do", Color.Blue);
            




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
   
            if(!string.IsNullOrEmpty(openFileDialog1.FileName))
            {
                StreamWriter escribir = new StreamWriter(openFileDialog1.FileName);
                foreach (object line in richTextBox1.Lines)
                {
                    escribir.WriteLine(line);
                }
                escribir.Close();
            }
            else if (!string.IsNullOrEmpty(saveFileDialog1.FileName))
            {
                StreamWriter escribir = new StreamWriter(saveFileDialog1.FileName);
                foreach (object line in richTextBox1.Lines)
                {
                    escribir.WriteLine(line);
                }
                escribir.Close();
            }
            else
            {
                saveFileDialog1.Filter = "Documento de texto|*.txt";
                saveFileDialog1.Title = "Documento de Texto";
                saveFileDialog1.FileName = "Sin titulo 1";
                var resultado = saveFileDialog1.ShowDialog();
                if (resultado == DialogResult.OK)
                {
                    rutaArchivo = "";
                    StreamWriter escribir = new StreamWriter(saveFileDialog1.FileName);
                    foreach (object line in richTextBox1.Lines)
                    {
                        escribir.WriteLine(line);
                    }
                    rutaArchivo = saveFileDialog1.FileName;
                    escribir.Close();
                }
            }

        }

        private void EjecutaAnalizadorLexico()
        {
         
           /* treeView1.Nodes.Add("Parent");
            treeView1.Nodes[0].Nodes.Add("1");
            treeView1.Nodes[0].Nodes.Add("2");
            treeView1.Nodes[0].Nodes[1].Nodes.Add("2");
            treeView1.Nodes[0].Nodes[1].Nodes[0].Nodes.Add("Great Grandchild");
            TreeNode buscado = treeView1.Nodes.Find("1", true)[0];
            treeView1.SelectedNode = buscado;
            treeView1.Nodes.Add("Hola");*/

            //Borro la informacion que hay en los archivos que utilizare para mostrar los resultados de la compilacion
             string ruta1 = "salida.txt";
             string ruta2 = "errores.txt";
             string ruta3 = "salidaSintactico.txt";
             string ruta4 = "erroresSintactico.txt";
             lblErrores.Text = "";
             lblSalida.Text = "";
             lblSalidaSintactico.Text = "";
            try
            {
                File.WriteAllText(ruta1, "");
                File.WriteAllText(ruta2, "");
                File.WriteAllText(ruta3, "");
                File.WriteAllText(ruta4, "");

            }
            catch (Exception err)
            {
                MessageBox.Show("A ocurrido un error");
            }
            //Compilo mi analizador lexico
            
            int Y = 10;
            try
            {
                string command = "go run main.go Scanner.go sintactico.go -archivo " + rutaArchivo;
                ExecuteCommand(command);

                //Mostramos los resultados del analisis
                string[] lines = System.IO.File.ReadAllLines("salida.txt");
                foreach (string line in lines)
                {
                    string[] resultado = line.Split(' ');

                    dataGridView1.Rows.Add(resultado[0], resultado[1], resultado[2]);
                    // Use a tab to indent each line of the file.
                    //lblSalida.Text += line + "\n";
                }

                //Mostramos los resultados de los Sintacticos
                string[] linesEr = System.IO.File.ReadAllLines("erroresSintactico.txt");
                foreach (string line in linesEr)
                {
                 // Use a tab to indent each line of the file.
                    lblErrores.Text += line + "\n";
                }

               //Mostramos los resultados del analisis Sintactico
  
                 string[] linesErSintactico = System.IO.File.ReadAllLines("salidaSintactico.txt");
                 foreach (string line in linesErSintactico)
                 {
                    
                   lblSalidaSintactico.Text += line + "\n";
                 }
                 
             }
             catch (Exception err)
             {
                 MessageBox.Show("Error no se ah podido compilar " + err.Message);
             }
        }

        

        private void toolStripButton7_Click(object sender, EventArgs e)
        {
      
            if(!string.IsNullOrEmpty(rutaArchivo))
            {
                dataGridView1.Rows.Clear();
                EjecutaAnalizadorLexico();
                
            }
            else
            {
                saveFileDialog1.Filter = "Documento de texto|*.txt";
                saveFileDialog1.Title = "Documento de Texto";
                saveFileDialog1.FileName = "Sin titulo 1";
                var resultado = saveFileDialog1.ShowDialog();
                if (resultado == DialogResult.OK)
                {
                    rutaArchivo = "";
                    StreamWriter escribir = new StreamWriter(saveFileDialog1.FileName);
                    foreach (object line in richTextBox1.Lines)
                    {
                        escribir.WriteLine(line);
                    }
                    rutaArchivo = saveFileDialog1.FileName;
                    escribir.Close();
                }
                EjecutaAnalizadorLexico();
            }
            

        }

        private void ExecuteCommand(string _Command)
        {
            //Indicamos que deseamos inicializar el proceso cmd.exe junto a un comando de arranque. 
            //(/C, le indicamos al proceso cmd que deseamos que cuando termine la tarea asignada se cierre el proceso).
            //Para mas informacion consulte la ayuda de la consola con cmd.exe /? 
            System.Diagnostics.ProcessStartInfo procStartInfo = new System.Diagnostics.ProcessStartInfo("cmd", "/c " + _Command);
            // Indicamos que la salida del proceso se redireccione en un Stream
            procStartInfo.RedirectStandardOutput = true;
            procStartInfo.UseShellExecute = false;
            //Indica que el proceso no despliegue una pantalla negra (El proceso se ejecuta en background)
            procStartInfo.CreateNoWindow = false;
            //Inicializa el proceso
            System.Diagnostics.Process proc = new System.Diagnostics.Process();
            proc.StartInfo = procStartInfo;
            proc.Start();
            //Consigue la salida de la Consola(Stream) y devuelve una cadena de texto
            string result = proc.StandardOutput.ReadToEnd();
            
            //Muestra en pantalla la salida del Comando
            //label1.Text = result.ToString();

        }

        private void abrirToolStripMenuItem_Click(object sender, EventArgs e)
        {
            //OpenFileDialog abrir = new OpenFileDialog();

            openFileDialog1.Filter = "Documento de texto|*.txt";
            openFileDialog1.Title = "Abrir";
            openFileDialog1.FileName = "Sin titulo 1";
            var resultado = openFileDialog1.ShowDialog();

            if (resultado == DialogResult.OK)
            {
                rutaArchivo = "";
                StreamReader leer = new StreamReader(openFileDialog1.FileName);
                richTextBox1.Text = leer.ReadToEnd();
                rutaArchivo = openFileDialog1.FileName;
                leer.Close();
                ScannerDoc();
                
            }
            

        }

        private void toolStripButton1_Click(object sender, EventArgs e)
        {
            
        }

        private void toolStripButton3_Click(object sender, EventArgs e)
        {
            openFileDialog1.Filter = "Documento de texto|*.txt";
            openFileDialog1.Title = "Abrir";
            openFileDialog1.FileName = "Sin titulo 1";
            var resultado = openFileDialog1.ShowDialog();

            if (resultado == DialogResult.OK)
            {
                rutaArchivo = "";
                StreamReader leer = new StreamReader(openFileDialog1.FileName);
                richTextBox1.Text = leer.ReadToEnd();
                rutaArchivo = openFileDialog1.FileName;
                leer.Close();
                ScannerDoc();
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
            saveFileDialog1.Filter = "Documento de texto|*.txt";
            saveFileDialog1.Title = "Documento de Texto";
            saveFileDialog1.FileName = "Sin titulo 1";
            var resultado = saveFileDialog1.ShowDialog();
            if (resultado == DialogResult.OK)
            {
                rutaArchivo = "";
                StreamWriter escribir = new StreamWriter(saveFileDialog1.FileName);
                foreach (object line in richTextBox1.Lines)
                {
                    escribir.WriteLine(line);
                }
                rutaArchivo = saveFileDialog1.FileName;
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
            //SaveFileDialog guardar = new SaveFileDialog();


            saveFileDialog1.Filter = "Documento de texto|*.txt";
            saveFileDialog1.Title = "Documento de Texto";
            saveFileDialog1.FileName = "Sin titulo 1";
            var resultado = saveFileDialog1.ShowDialog();
            if (resultado == DialogResult.OK)
            {
                rutaArchivo = "";
                StreamWriter escribir = new StreamWriter(saveFileDialog1.FileName);
                foreach (object line in richTextBox1.Lines)
                {
                    escribir.WriteLine(line);
                }
                rutaArchivo = saveFileDialog1.FileName;
                escribir.Close();
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
                //ScannerDoc();
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

        private void openFileDialog1_FileOk(object sender, CancelEventArgs e)
        {

        }

        private void toolStripButton5_Click(object sender, EventArgs e)
        {
            
        }

        private void tabPage2_Click(object sender, EventArgs e)
        {

        }

        private void label2_Click(object sender, EventArgs e)
        {

        }

        private void richTextBox1_TextChanged(object sender, EventArgs e)
        {

        }
    }
}
